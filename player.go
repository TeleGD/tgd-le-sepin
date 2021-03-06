package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func (g *Game) movePlayer() {
	// Check user inputs
	for _, key := range g.keys {
		if key == ebiten.KeyArrowUp {
			user.Position = 0
			break
		} else if key == ebiten.KeyArrowLeft {
			user.Position = 3 * math.Pi / 2
			break
		} else if key == ebiten.KeyArrowRight {
			user.Position = math.Pi / 2
			break
		} else if key == ebiten.KeyArrowDown {
			user.Position = math.Pi
			break
		}
	}
}

func (g *Game) drawPlayer(screen *ebiten.Image) {
	w, h := player.Size()
	op := &ebiten.DrawImageOptions{}

	// Handle player image
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Rotate(user.Position)
	op.GeoM.Translate(screenWidth/2, screenHeight/2)
	screen.DrawImage(player, op)
	text.Draw(screen, fmt.Sprintf("Score : %d", -user.Score), text.FaceWithLineHeight(mplusBigFont, 80), 30, 30, color.White)
	text.Draw(screen, fmt.Sprintf("FPS : %0.2f", ebiten.CurrentFPS()), text.FaceWithLineHeight(mplusBigFont, 80), screenWidth-200, 30, color.White)
	text.Draw(screen, "Goal : Reach 250 !", text.FaceWithLineHeight(mplusBigFont, 80), 30, 80, color.White)
	if -user.Score >= 250 || user.HitScore {
		user.HitScore = true
		text.Draw(screen, fmt.Sprintf("Flag : %s", genFlag()), text.FaceWithLineHeight(mplusBigFont, 80), 30, 130, color.White)
	}
}

func genFlag() (decryptedString string) {
	enc, _ := hex.DecodeString("73367639792442264529482B4D6251655468576D5A7134743777217A25432A46")
	key, _ := hex.DecodeString("e754c3db4c4f0190a03b3d99732442791a27cc56a8abf631b1f7b05a0a22b05c2923cdef4b4e")
	block, err := aes.NewCipher(enc)
	if err != nil {
		panic(err.Error())
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := aesGCM.NonceSize()
	nonce, ciphertext := key[:nonceSize], key[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return fmt.Sprintf("%s", plaintext)
}
