package dependencia

import "github.com/francoganga/go_reno2/pkg/domain/user"

type AreaSudocu string
type Dependencia struct {
	Nombre      string
	AreaSudocu  AreaSudocu
	Autorizante *user.User
}

func New(nombre string, areaSudocu string, autorizante *user.User) *Dependencia {
    return &Dependencia{
        Nombre: nombre,
        AreaSudocu: AreaSudocu(areaSudocu),
        Autorizante: autorizante,
    }
}
