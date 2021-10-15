package main

import (
	"bytes"
	"image"
	"log"
	"time"

	_ "embed"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	screenWidth  = 1000
	screenHeight = 800
)

var (
	player          *ebiten.Image
	ennemies        *ebiten.Image
	spawnedEnnemies []*Ennemy
	playerPos       float64
	playerScore     int
	lastEnnemy      time.Time
)

type Game struct {
	count int
	keys  []ebiten.Key

	// Relative to audio
	player       *audio.Player
	audioContext *audio.Context
}

const (
	sampleRate = 22255

	introLengthInSecond = 5
	loopLengthInSecond  = 4
)

func (g *Game) Update() error {
	g.count++
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])

	if g.player != nil {
		return nil
	}

	if g.audioContext == nil {
		g.audioContext = audio.NewContext(sampleRate)
	}

	// Decode an Ogg file.
	// oggS is a decoded io.ReadCloser and io.Seeker.
	oggS, err := vorbis.Decode(g.audioContext, bytes.NewReader(audioLoop))
	if err != nil {
		return err
	}

	// Create an infinite loop stream from the decoded bytes.
	// s is still an io.ReadCloser and io.Seeker.
	s := audio.NewInfiniteLoopWithIntro(oggS, introLengthInSecond*4*sampleRate, loopLengthInSecond*4*sampleRate)

	g.player, err = g.audioContext.NewPlayer(s)
	if err != nil {
		return err
	}

	g.player.SetVolume(0.1)
	// Play the infinite-length stream. This never ends.
	g.player.Play()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.pickEnnemy()
	g.drawPlayer(screen)
	g.drawAllEnnemies(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

//go:embed assets/player.png
var playerBytes []byte

//go:embed assets/enemy.png
var ennemyBytes []byte

//go:embed assets/loop.ogg
var audioLoop []byte

//go:embed assets/hit.ogg
var hitSound []byte

var (
	mplusNormalFont font.Face
	mplusBigFont    font.Face
)

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	mplusBigFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    32,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	img, _, err := image.Decode(bytes.NewReader(playerBytes))
	if err != nil {
		log.Fatal(err)
	}
	player = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(ennemyBytes))
	if err != nil {
		log.Fatal(err)
	}
	ennemies = ebiten.NewImageFromImage(img)

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Le sépan")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
