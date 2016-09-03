package main

var currentID int

var dataObjs EncryptedData

// RepoFindObj gets the object with given ID from the database
func RepoFindObj(ID int) Encrypted {

	for _, t := range dataObjs {
		if t.ID == ID {
			return t
		}
	}
	// return empty Obj if not found
	return Encrypted{}
}

// RepoCreateObj manages the ID of the object and then stores the object
func RepoCreateObj(t Encrypted) Encrypted {

	//TODO: If the ID already exists, change it so it doesn't collide
	//----> Implement idAlreadyExists

	if t.ID == 0 {
		currentID++
		t.ID = currentID
	}

	dataObjs = append(dataObjs, t)
	return t
}

func idAlreadyExists(id int) bool {
	for _, t := range dataObjs {
		if t.ID == id {
			return true
		}
	}
	return false
}
