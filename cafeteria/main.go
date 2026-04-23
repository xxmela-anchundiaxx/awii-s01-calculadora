func main() {

	var clientes []Cliente
	var productos []Producto
	var pedidos []Pedido

	var opcion int

	for {
		fmt.Println("\n===== SISTEMA DE CAFETERÍA =====")
		fmt.Println("1. Gestionar Clientes")
		fmt.Println("2. Gestionar Productos")
		fmt.Println("3. Registrar Pedido")
		fmt.Println("4. Salir")
		fmt.Print("Seleccione una opción: ")
		fmt.Scan(&opcion)

		switch opcion {

		// -------- CLIENTES --------
		case 1:
			var opCliente int
			fmt.Println("\n--- CLIENTES ---")
			fmt.Println("1. Agregar")
			fmt.Println("2. Listar")
			fmt.Println("3. Eliminar")
			fmt.Print("Opción: ")
			fmt.Scan(&opCliente)

			switch opCliente {
			case 1:
				var id int
				var nombre, carrera string
				var saldo float64

				fmt.Print("ID: ")
				fmt.Scan(&id)
				fmt.Print("Nombre: ")
				fmt.Scan(&nombre)
				fmt.Print("Carrera: ")
				fmt.Scan(&carrera)
				fmt.Print("Saldo: ")
				fmt.Scan(&saldo)

				clientes = AgregarClientes(clientes, Cliente{id, nombre, carrera, saldo})

			case 2:
				ListarClientes(clientes)

			case 3:
				var id int
				fmt.Print("ID a eliminar: ")
				fmt.Scan(&id)
				clientes = EliminarCliente(clientes, id)
			}

		// -------- PRODUCTOS --------
		case 2:
			var opProducto int
			fmt.Println("\n--- PRODUCTOS ---")
			fmt.Println("1. Agregar")
			fmt.Println("2. Listar")
			fmt.Println("3. Eliminar")
			fmt.Print("Opción: ")
			fmt.Scan(&opProducto)

			switch opProducto {
			case 1:
				var id int
				var nombre, categoria string
				var precio float64
				var stock int

				fmt.Print("ID: ")
				fmt.Scan(&id)
				fmt.Print("Nombre: ")
				fmt.Scan(&nombre)
				fmt.Print("Precio: ")
				fmt.Scan(&precio)
				fmt.Print("Stock: ")
				fmt.Scan(&stock)
				fmt.Print("Categoría: ")
				fmt.Scan(&categoria)

				productos = AgregarProductos(productos, Producto{id, nombre, precio, stock, categoria})

			case 2:
				ListarProductos(productos)

			case 3:
				var id int
				fmt.Print("ID a eliminar: ")
				fmt.Scan(&id)
				productos = EliminarProducto(productos, id)
			}

		// -------- PEDIDOS --------
		case 3:
			var clienteID, productoID, cantidad int
			var fecha string

			fmt.Print("ID Cliente: ")
			fmt.Scan(&clienteID)
			fmt.Print("ID Producto: ")
			fmt.Scan(&productoID)
			fmt.Print("Cantidad: ")
			fmt.Scan(&cantidad)
			fmt.Print("Fecha: ")
			fmt.Scan(&fecha)

			var err error
			pedidos, err = RegistrarPedido(clientes, productos, pedidos, clienteID, productoID, cantidad, fecha)

			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Pedido registrado correctamente")
			}

		case 4:
			fmt.Println("Saliendo del sistema...")
			return

		default:
			fmt.Println("Opción inválida")
		}
	}
}
