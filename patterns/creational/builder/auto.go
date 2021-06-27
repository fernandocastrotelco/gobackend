package creational

type auto struct {
	ruedas  int
	puertas int
	marca   string
	modelo  string
	anio    int
}

type AutoBuilder interface {
	SetRuedas(int) AutoBuilder
	SetPuertas(int) AutoBuilder
	SetMarca(string) AutoBuilder
	SetModelo(string) AutoBuilder
	SetAnio(int) AutoBuilder
	Build() auto
}

type Builder struct {
	a auto
}

func (b *Builder) SetRuedas(r int) AutoBuilder {
	b.a.ruedas = r
	return b
}

func (b *Builder) SetPuertas(p int) AutoBuilder {
	b.a.puertas = p
	return b
}

func (b *Builder) SetMarca(m string) AutoBuilder {
	b.a.marca = m
	return b
}

func (b *Builder) SetModelo(m string) AutoBuilder {
	b.a.modelo = m
	return b
}

func (b *Builder) SetAnio(n int) AutoBuilder {
	b.a.anio = n
	return b
}

func (b *Builder) Build() auto {
	return b.a
}

// type MetodoCualquiera struct {
// 	builder Builder
// }

// func (m *MetodoCualquiera) construirAuto() {
// 	nuevoAuto := m.builder

// 	nuevoAuto.SetAnio(2020)

// 	nuevoAuto.SetRuedas(4).SetPuertas(5)

// 	nuevoAuto.SetMarca("Ford")

// 	fordFocus := nuevoAuto.SetModelo("Focus").Build()

// 	println(fordFocus.modelo)
// }
