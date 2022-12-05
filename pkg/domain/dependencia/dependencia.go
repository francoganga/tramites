package dependencia

import "github.com/francoganga/go_reno2/pkg/domain/user"

type AreaSudocu string
type Dependencia struct {
	Nombre      string
	AreaSudocu  AreaSudocu
	Autorizante *user.User
}
