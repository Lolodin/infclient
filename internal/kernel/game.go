package kernel

import (
	"github.com/Lolodin/infclient/fonts"
	"github.com/Lolodin/infclient/internal/component"
	"github.com/Lolodin/infclient/internal/entity"
	"github.com/Lolodin/infclient/internal/lang"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/f64"
	"golang.org/x/text/language"
	"image/color"
	"log"
)

type GameSystem interface {
	GetComponents() []string
	AddEntity(*entity.Entity, int)
	Remove(*entity.Entity) bool
	Update(*State)
	Draw(*State, *ebiten.Image)
	Load(*State)
	Exit(*State)
	Enter(*State)
}

type gameInteractive interface {
	Name() string
	Load(*State)
	Exit(*State)
	Enter(*State)
	Update(*State)
	Draw(*State, *ebiten.Image)
}
type Options struct {
	RenderWidth, RenderHeight int
	Title                     string
}

type Auth interface {
	SetLogin(str string)
	SetPassword(str string)
	SetToken(str string)
	GetLogin() string
	GetToken() string
	GetPassword() string
}

type State struct {
	Options
	worlds       map[string]gameInteractive
	activeWorlds []string
	Events       []interface{}
	Camera       *Camera

	Controls    map[component.Control]*component.InputData
	MouseInputs map[ebiten.MouseButton]component.Control
	//KeyInputs        map[ebiten.Key]component.Control
	CursorX, CursorY float64

	IsCursorHovering bool

	Fonts map[string]font.Face

	Colors map[string]color.NRGBA
	Lang   *lang.Translator
	Auth   Auth
}

func NewState(options Options, tag language.Tag) *State {
	s := &State{}
	s.Camera = &Camera{ViewPort: f64.Vec2{float64(options.RenderWidth), float64(options.RenderHeight)}}
	s.worlds = map[string]gameInteractive{}
	s.Controls = map[component.Control]*component.InputData{}
	s.RenderWidth = options.RenderWidth
	s.RenderHeight = options.RenderHeight
	s.Auth = &AuthData{}
	s.Title = options.Title
	s.Lang = lang.NewTranslator()
	s.Lang.Lang = tag
	tt, err := opentype.Parse(fonts.Alundra)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	s.Fonts = map[string]font.Face{"std": mplusNormalFont}
	//Add config with lang
	s.Lang.AddLocalizer("langs/RU.json", language.Russian)
	ebiten.SetWindowSize(s.RenderWidth*2, s.RenderHeight*2)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	ebiten.SetWindowTitle(s.Title)

	return s
}

func (s *State) Update() error {
	s.UpdateCursor()

	for _, worldName := range s.activeWorlds {
		s.worlds[worldName].Update(s)
	}
	return nil
}

func (s *State) Draw(screen *ebiten.Image) {
	for _, worldName := range s.activeWorlds {
		s.worlds[worldName].Draw(s, screen)
	}
}

func (s *State) Layout(_, _ int) (int, int) {
	return s.RenderWidth, s.RenderHeight
}

// Map the controls to the input fields
func (s *State) UpdateControls() {
	// Clear any existing values
	s.Controls = map[component.Control]*component.InputData{}

	for _, control := range s.MouseInputs {
		s.Controls[control] = nil
	}
}

func (s *State) LoadWorld(w gameInteractive) {
	w.Load(s)
	s.worlds[w.Name()] = w
}

// Build slices for exiting and entering worlds based on what
// worlds are currently active and those that will be. Then
// exit and enter all of those worlds.
func (s *State) ActivateWorlds(names ...string) {
	exitingWorlds := sliceStringDifference(s.activeWorlds, names)
	enteringWorlds := sliceStringDifference(names, s.activeWorlds)

	for _, worldName := range exitingWorlds {
		s.worlds[worldName].Exit(s)
	}

	s.activeWorlds = names

	for _, worldName := range enteringWorlds {
		s.worlds[worldName].Enter(s)
	}
}

func (s *State) UpdateCursor() {
	x, y := ebiten.CursorPosition()
	s.CursorX = float64(x)
	s.CursorY = float64(y)
}

func (s *State) SetControl(control component.Control, data *component.InputData) {
	s.Controls[control] = data
}

func (s *State) AddEvent(event interface{}) {
	s.Events = append(s.Events, event)
}

func (s *State) ClearEvents() {
	s.Events = []interface{}{}
}

func (s *State) SetIsCursorHovering(isHovering bool) {
	s.IsCursorHovering = isHovering
}

// Find the unique strings in the first slice between 2 slices
// of strings
func sliceStringDifference(a, b []string) []string {
	uniqueStrings := []string{}
	isStringUnique := true

	for _, stringA := range a {
		for _, stringB := range b {
			if stringA == stringB {
				isStringUnique = false
			}
		}

		if isStringUnique {
			uniqueStrings = append(uniqueStrings, stringA)
		}

		isStringUnique = true
	}

	return uniqueStrings
}
