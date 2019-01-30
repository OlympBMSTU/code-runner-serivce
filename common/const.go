package common

const (
	FPCCOMPIL  = "fpc"
	GCCCOMPIL  = "gcc"
	GPPCOMPIL  = "g++"
	PY27INTERP = "python2.7"
	PY3INTERP  = "python3"
)

const (
	PASCAL = "pascal"
	CLANG  = "c"
	CPP    = "cpp"
	PY27   = "pyhton2.7"
	PY3    = "python3"
)

var compilers = map[string]string{
	PASCAL: FPCCOMPIL,
	CLANG:  GCCCOMPIL,
	CPP:    GPPCOMPIL,
	PY27:   PY27INTERP,
	PY3:    PY3INTERP,
}
