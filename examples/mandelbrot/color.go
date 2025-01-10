package main

/*
https://stackoverflow.com/questions/2353211/hsl-to-rgb-color-conversion
h [0, 1)
s [0, 1)
l [0, 1)
*/
func hslToRGB(h, s, l float64) (int, int, int) {
	var r, g, b float64
	if s == 0 {
		r, g, b = l, l, l
	} else {
		var q, p float64
		if l < 0.5 {
			q = l * (1 + s)
		} else {
			q = l + s - l*s
		}
		p = 2*l - q
		r = hueToRGB(p, q, h+1.0/3.0)
		g = hueToRGB(p, q, h)
		b = hueToRGB(p, q, h-1.0/3.0)

		// if r > 1 || g > 1 || b > 1 {
		// 	fmt.Println("there is a problem")
		// }
		// if r < 0 || g < 0 || b < 0 {
		// 	fmt.Println("there is a problem 2")
		// }
	}
	return int(r * 255), int(g * 255), int(b * 255)
}

func hueToRGB(p, q, t float64) float64 {
	if t < 0 {
		t += 1
	}
	if t > 1 {
		t -= 1
	}
	switch {
	case t < 1.0/6.0:
		return p + (q-p)*6*t
	case t < 1.0/2.0:
		return q
	case t < 2.0/3.0:
		return p + (q-p)*(2.0/3.0-t)*6
	default:
		return p
	}
}
