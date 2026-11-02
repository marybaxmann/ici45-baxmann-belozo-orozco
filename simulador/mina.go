package main

type TipoNodo string
type TipoTramo string

type Nodo struct {
    ID        string
    Tipo      TipoNodo
    Capacidad int
}

type Tramo struct {
    ID        string
    Origen    string
    Destino   string
    Tipo      TipoTramo
    Tiempo    int
}