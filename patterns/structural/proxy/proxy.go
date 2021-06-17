package main

import "fmt"

func main() {
	fmt.Println("Implementando proxy pattern")
}

// vamos a implementar proxy protegiendo una base de datos con una cache
type User struct {
	ID int32
}

// UserFinder es la interfaz que van a implementar tanto la base como el proxy
type UserFinder interface {
	FindUser(id int32) (User, error)
}

// UserList es el objeto q va a implementar la interfaz
type UserList []User

// chequea si el type UserList puede implementar la interfaz UserFinder
var _ UserFinder = (*UserList)(nil)

// chequea si el type UserListProxy puede implementar la interfaz UserFinder
var _ UserFinder = (*UserListProxy)(nil)

//FindUser va a recorrer la lista buscando un user con el mismo
//id que el parametro o devuelve un error si no lo encuentra
func (t *UserList) FindUser(id int32) (User, error) {
	for i := 0; i < len(*t); i++ {
		if (*t)[i].ID == id {
			return (*t)[i], nil
		}
	}

	return User{}, fmt.Errorf("user %d could not be found", id)
}

//AddUser agrega un nuevo usuario al final de la lista
func (t *UserList) addUser(newUser User) {
	*t = append(*t, newUser)
}

// UserListProxy es el objeto que va a actuar de proxy
type UserListProxy struct {
	MockedDatabase *UserList
	StackCache     UserList
	// el tamaño maximo q va a cachear el proxy
	StackSize int
	// ayuda a saber si la ultima busqueda uso cache
	LastSearchUsedCache bool
}

// addUserToStack recibe un usuario y lo agrega al stack de cache.
// si el stack esta lleno quita el primer elemento antes de agregar el nuevo.
func (u *UserListProxy) addUserToStack(user User) {
	if len(u.StackCache) >= u.StackSize {
		u.StackCache = append(u.StackCache[1:], user)
	} else {
		u.StackCache.addUser(user)
	}
}

//FindUser will search for the specified name in the parameter in the cache
//list. If it finds it, it will return it. If not, it will search in the heavy
//list. Finally, if it's not in the heavy list, it will return an error
//(generated from the heavy list)
func (u *UserListProxy) FindUser(id int32) (User, error) {
	//Search for the object in the cache list first
	user, err := u.StackCache.FindUser(id)
	if err == nil {
		fmt.Println("Returning user from cache")
		u.LastSearchUsedCache = true
		return user, nil
	}

	//Object is not in the cache list. Search in the heavy list
	user, err = u.MockedDatabase.FindUser(id)
	if err != nil {
		return User{}, err
	}

	//Adds the new user to the stack, removing the last if necessary
	u.addUserToStack(user)

	fmt.Println("Returning user from database")
	u.LastSearchUsedCache = false
	return user, nil
}
