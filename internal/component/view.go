package component

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"io"
	"os"
	"strings"
)

//Для отрисовки компонента
type ViewComponent struct {
	Image                      *ebiten.Image
	Frame, Duration            int
	Sequence, PreviousSequence string
	Time                       float64
	Frames                     []*image.Rectangle
	Sequences                  map[string]*ViewSequence
	IsDraw                     bool
}

type ViewSequence struct {
	From, To   int
	ShouldLoop bool
	Direction  string
}

// Загружает ресурсы для отображения компонента
// Если анимация не нужна, второй компонент можно оставить пустым
func NewViewComponent(imagePath, animationPath string) (*ViewComponent, error) {
	c := &ViewComponent{
		IsDraw:    true,
		Sequences: map[string]*ViewSequence{},
	}

	// Загрузка картинок
	file, err := os.Open(imagePath)
	defer file.Close()
	if err != nil {
		return nil, fmt.Errorf("loading '%s' image: %s", imagePath, err)
	}
	buff, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("loading '%s' image: %s", imagePath, err)
	}
	rawImg, _, err := image.Decode(bytes.NewReader(buff))
	if err != nil {
		return c, fmt.Errorf("decoding appearance image: %s", err)
	}
	c.Image = ebiten.NewImageFromImage(rawImg)

	// Без Анимации
	if animationPath == "" {
		w, h := c.Image.Size()
		c.Frames = []*image.Rectangle{{
			Min: image.Point{0, 0},
			Max: image.Point{w, h},
		}}
		return c, nil
	}

	// Логика Анимации
	anim, err := newAnimationFromFile(animationPath)
	if err != nil {
		return c, fmt.Errorf("loading appearance animation: %s", err)
	}
	for _, f := range anim.Frames {
		c.Duration += f.Duration
		rect := image.Rect(
			f.Frame.X, f.Frame.Y,
			f.Frame.X+f.Frame.W, f.Frame.Y+f.Frame.H,
		)
		c.Frames = append(c.Frames, &rect)
	}
	for _, t := range anim.Meta.FrameTags {
		// Логика для работы с Aseprite
		nameSections := strings.Split(t.Name, "_")
		shouldLoop := nameSections[len(nameSections)-1] == "loop"
		name := nameSections[0]
		c.Sequences[name] = &ViewSequence{
			From:       t.From,
			To:         t.To,
			ShouldLoop: shouldLoop,
			Direction:  t.Direction,
		}
		// Sequence as the default
		if c.Sequence == "" {
			c.Sequence = name
		}
	}

	return c, nil
}

// Struct to match Aseprite JSON output
type animation struct {
	Frames []struct {
		Duration int `json:"duration"`
		Frame    struct {
			X int `json:"x"`
			Y int `json:"y"`
			W int `json:"w"`
			H int `json:"h"`
		} `json:"frame"`
	} `json:"frames"`
	Meta struct {
		Size struct {
			W int `json:"w"`
			H int `json:"h"`
		} `json:"size"`
		FrameTags []struct {
			Name      string `json:"name"`
			From      int    `json:"from"`
			To        int    `json:"to"`
			Direction string `json:"direction"`
		} `json:"frameTags"`
	} `json:"meta"`
}

// Load an Aseprite JSON file as an animation
func newAnimationFromFile(path string) (*animation, error) {
	ofile, err := os.Open(path)
	defer ofile.Close()
	if err != nil {
		return nil, fmt.Errorf("loading '%s' file: %s", path, err)
	}
	file, err := io.ReadAll(ofile)
	if err != nil {
		return nil, fmt.Errorf("loading animation json: %s", err)
	}
	var anim animation
	err = json.Unmarshal(file, &anim)
	if err != nil {
		return &anim, fmt.Errorf("unmarshalling animation json: %s", err)
	}
	return &anim, nil
}
