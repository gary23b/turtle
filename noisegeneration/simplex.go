package noisegeneration

import "github.com/ojrac/opensimplex-go"

// https://dbriemann.github.io/blog/7-making-a-game-with-go-and-pixel-2-procedural-content-generation-pcg.html
// https://github.com/ojrac/opensimplex-go
type Simplex32LayerNoise2D struct {
	noise opensimplex.Noise32
	seed  int64

	layers      int
	persistence float32
	freq        float32
	low         float32
	high        float32
}

func NewSimplex32LayerNoise2D(seed int64, layers int, persistence, freq, low, high float32) *Simplex32LayerNoise2D {
	ret := &Simplex32LayerNoise2D{}

	ret.noise = opensimplex.New32(seed)
	ret.seed = seed
	ret.layers = layers
	ret.persistence = persistence
	ret.freq = freq
	ret.low = low
	ret.high = high

	return ret
}

func (s *Simplex32LayerNoise2D) LayerNoise(x, y float32) (result float32) {
	var ampSum float32 = 0.0
	var amp float32 = 1.0
	freq := s.freq

	for i := 0; i < s.layers; i++ {
		result += s.noise.Eval2(x*freq, y*freq) * amp
		ampSum += amp
		amp *= s.persistence
		freq *= 2
	}

	result /= ampSum

	result = result*(s.high-s.low)/2 + (s.high+s.low)/2
	return
}
