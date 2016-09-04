package main

var currentID int

var dataObjs EncryptedData

// RepoFindObj gets the object with given ID from the database
func RepoFindObj(ID int) Encrypted {

	for _, enc := range dataObjs {
		if enc.ID == ID {
			return enc
		}
	}
	// return empty Obj if not found
	return Encrypted{}
}

// RepoCreateObj manages the ID of the object and then stores the object
func RepoCreateObj(enc Encrypted) Encrypted {

	//TODO: If the ID already exists, change it so it doesn't collide
	//----> Implement idAlreadyExists

	if enc.ID == 0 {
		currentID++
		enc.ID = currentID
	}

	dataObjs = append(dataObjs, enc)
	return enc
}

func idAlreadyExists(id int) bool {
	for _, enc := range dataObjs {
		if enc.ID == id {
			return true
		}
	}
	return false
}
