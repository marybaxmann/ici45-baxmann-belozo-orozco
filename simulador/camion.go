package main

type EstadoCamion string

const (
	ViajandoVacio   EstadoCamion = "ViajandoVacio"
	EnColaPala      EstadoCamion = "EnColaPala"
	Cargando        EstadoCamion = "Cargando"
	Transportando   EstadoCamion = "Transportando"
	EnColaStockpile EstadoCamion = "EnColaStockpile"
	Descargando     EstadoCamion = "Descargando"
	EnFalla         EstadoCamion = "EnFalla"
)

type Camion struct {
	ID        string
	Capacidad float64
	Estado    EstadoCamion
	Ubicacion string
}

func NuevoCamion(id string, cap float64) *Camion {
	return &Camion{
		ID:        id,
		Capacidad: cap,
		Estado:    ViajandoVacio,
		Ubicacion: "E1",
	}
}
