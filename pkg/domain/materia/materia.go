package materia

import (
	"bytes"
	"strconv"
)

type InstitutoNombre string
type MateriaNombre string
type PropuestaNombre string
type InstitutoId int
type PropuestaId int
type MateriaId int

type Materia struct {
	Nombre      MateriaNombre
	Instituto   InstitutoNombre
	Propuesta   PropuestaNombre
	MateriaId   MateriaId
	InstitutoId InstitutoId
	PropuestaId PropuestaId
}

func (m *Materia) String() string {
	var out bytes.Buffer

	out.WriteString("{")
	out.WriteString("nombre: ")
	out.WriteString(string(m.Nombre))
	out.WriteString(", instituto: ")
	out.WriteString(string(m.Instituto))
	out.WriteString(", propuesta: ")
	out.WriteString(string(m.Propuesta))

	out.WriteString(", materiaId: ")
	out.WriteString(strconv.Itoa(int(m.MateriaId)))

	out.WriteString(", institutoId: ")
	out.WriteString(strconv.Itoa(int(m.InstitutoId)))

	out.WriteString(", PropuestaId: ")
	out.WriteString(strconv.Itoa(int(m.PropuestaId)))

	out.WriteString("}")

	return out.String()
}
