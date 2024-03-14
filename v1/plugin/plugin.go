package plugin

func (p Type) PathName() string {
	switch p {
	case Type_KIT:
		return "kit"
	case Type_EXECFILE:
		return "exec"
	case Type_CONTAINER:
		return "image"
	default:
		return "unknown"
	}
}
