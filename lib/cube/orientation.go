package cube

type Orientation uint16

const (
	RotateSquare       Orientation = 4
	RotateMirrorSquare             = 8
	RotateCube                     = 24
)

func (o Orientation) Tr(x, y, z int) (int, int, int) {
	o = o % 24
	switch o {
	// Facing negative Z
	case 1:
		return y, -x, z
	case 2:
		return -x, -y, z
	case 3:
		return -y, x, z
	// Facing positive Z
	case 4:
		return -x, y, -z
	case 5:
		return y, x, -z
	case 6:
		return x, -y, -z
	case 7:
		return -y, -x, -z
	// Facing negative Y
	case 8:
		return -x, z, y
	case 9:
		return -z, -x, y
	case 10:
		return x, -z, y
	case 11:
		return z, x, y
	// Facing positive Y
	case 12:
		return -x, -z, -y
	case 13:
		return z, -x, -y
	case 14:
		return x, z, -y
	case 15:
		return -z, x, -y
	// Facing negative X
	case 16:
		return z, -y, x
	case 17:
		return y, z, x
	case 18:
		return -z, y, x
	case 19:
		return -y, -z, x
	// Facing positive X
	case 20:
		return -z, -y, -x
	case 21:
		return y, -z, -x
	case 22:
		return z, y, -x
	case 23:
		return -y, z, -x
	default:
		return x, y, z
	}
}
