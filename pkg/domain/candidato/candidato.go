package candidato

import "github.com/google/uuid"

type FilePath string

type CandidatoSubioArchivos struct {
	ID       uuid.UUID  `json:"id"`
	Archivos []FilePath `json:"archivos"`
}

type CandidatoNombre string
type Apellido string
type CandidatoEmail string

type Candidato struct {
	Nombre   CandidatoNombre
	Apellido Apellido
	Email    CandidatoEmail
	Archivos []FilePath
}

func New(nombre CandidatoNombre, apellido Apellido, email CandidatoEmail) *Candidato {
    return &Candidato{
        Nombre: nombre,
        Apellido: apellido,
        Email: email,
        Archivos: make([]FilePath, 0),
    }
}

func (c *Candidato) AddFile(path FilePath) {
    c.Archivos = append(c.Archivos, path)
}
