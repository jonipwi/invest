package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Function to compare two bits and return the result (1 if at least one bit is 1)
func bitwiseOR(bit1, bit2 int) int {
	if bit1 == 1 || bit2 == 1 {
		return 1
	}
	return 0
}

func bitwiseAND(bit1, bit2 int) int {
	if bit1 == 1 && bit2 == 1 {
		return 1
	}
	return 0
}

// Function to compare two bits using XOR and return the result (1 if bits are different, 0 if they are the same)
func bitwiseXOR(bit1, bit2 int) int {
	if bit1 != bit2 {
		return 1
	}
	return 0
}

// searchByWord searches the database for all rows with a specific word and returns a list of matching rows.
func searchByWord(db *Database, searchWord string) []*Word {
	// Assuming that the 'db' parameter is a slice of Word objects.

	// Create an empty list to store matching rows.
	matchingRows := []*Word{}

	// Iterate through the database and find rows with the specified word.
	for _, word := range db.Words {
		if word.Word == searchWord {
			// Create a new instance of Word for each matching row and append it to matchingRows.
			matchingRows = append(matchingRows, &Word{
				ID:     word.ID,
				Word:   word.Word,
				Weight: word.Weight,
			})
		}
	}

	// Return the list of matching rows.
	return matchingRows
}

// searchByID searches the database for all rows with a specific ID and returns a list of matching rows.
func searchByID(db *Database, id int) []*Word {
	// Assuming that the 'db' parameter is a slice of Word objects.

	// Create an empty list to store matching rows.
	matchingRows := []*Word{}

	// Iterate through the database and find rows with the specified ID.
	for _, word := range db.Words {
		if word.ID == id {
			// Create a new instance of Word for each matching row and append it to matchingRows.
			matchingRows = append(matchingRows, &Word{
				ID:     word.ID,
				Word:   word.Word,
				Weight: word.Weight,
			})
		}
	}

	// Return the list of matching rows.
	return matchingRows
}

func bitwiseORWords(word1, word2 string, db *Database) (int, string, error) {
	var weight1, weight2 int
	var id1, id2 int
	id1 = 0
	id2 = 0

	wordToSearch := word1
	matchingRowsByWord := searchByWord(db, wordToSearch)
	fmt.Println("Matching rows with word", wordToSearch, ":")
	for _, row1 := range matchingRowsByWord {
		//fmt.Println(row)
		idToSearch := row1.ID
		matchingRowsByID := searchByID(db, idToSearch)
		//fmt.Println("Matching rows with ID", idToSearch, ":")
		for _, row2 := range matchingRowsByID {
			//fmt.Println(row2)
			if row2.Word == word2 {
				id1 = row1.ID
				weight1 = row1.Weight
				id2 = row2.ID
				weight2 = row2.Weight
			}
		}
	}

	result := (weight1 == 1) || (weight2 == 1)
	if result == true {
		if weight1 == 1 {
			return id1, word1, nil
		} else if weight2 == 1 {
			return id2, word2, nil
		}
	}

	return 0, "", fmt.Errorf("no weight equal to 1")
}

func bitwiseANDWords(word1, word2 string, db *Database) (int, string, error) {
	var weight1, weight2 int
	var id1, id2 int
	id1 = 0
	id2 = 0

	wordToSearch := word1
	matchingRowsByWord := searchByWord(db, wordToSearch)
	fmt.Println("Matching rows with word", wordToSearch, ":")
	for _, row1 := range matchingRowsByWord {
		//fmt.Println(row)
		idToSearch := row1.ID
		matchingRowsByID := searchByID(db, idToSearch)
		//fmt.Println("Matching rows with ID", idToSearch, ":")
		for _, row2 := range matchingRowsByID {
			//fmt.Println(row2)
			if row2.Word == word2 {
				id1 = row1.ID
				weight1 = row1.Weight
				id2 = row2.ID
				weight2 = row2.Weight
			}
		}
	}

	result := (weight1 == 1) && (weight2 == 1)
	if result == true {
		if weight1 == 1 {
			return id1, word1, nil
		} else if weight2 == 1 {
			return id2, word2, nil
		}
	}

	return 0, "", fmt.Errorf("no weight equal to 1")
}

func bitwiseXORWords(word1, word2 string, db *Database) (int, string, error) {
	var weight1, weight2 int
	var id1, id2 int
	id1 = 0
	id2 = 0

	wordToSearch := word1
	matchingRowsByWord := searchByWord(db, wordToSearch)
	fmt.Println("Matching rows with word", wordToSearch, ":")
	for _, row1 := range matchingRowsByWord {
		//fmt.Println(row)
		idToSearch := row1.ID
		matchingRowsByID := searchByID(db, idToSearch)
		//fmt.Println("Matching rows with ID", idToSearch, ":")
		for _, row2 := range matchingRowsByID {
			//fmt.Println(row2)
			if row2.Word == word2 {
				id1 = row1.ID
				weight1 = row1.Weight
				id2 = row2.ID
				weight2 = row2.Weight
			}
		}
	}

	result := weight1 != weight2
	if result == true {
		if weight1 == 1 {
			return id1, word1, nil
		} else if weight2 == 1 {
			return id2, word2, nil
		}
	}

	return 0, "", fmt.Errorf("no weight equal to 1")
}

const (
	dbPath = "./data.json" // Change this path as per your preference
)

// Word represents a row in the database.
type Word struct {
	ID     int    `json:"id"`
	Word   string `json:"word"`
	Weight int    `json:"weight"`
}

// Database represents the whole JSON database.
type Database struct {
	Words []Word `json:"words"`
}

func main() {
	// Initialize the database
	db := &Database{}

	// Load data from the JSON file (if it exists)
	if err := loadDataFromJSON(dbPath, db); err != nil {
		log.Fatal(err)
	}

	// Insert some sample data into the database (if it doesn't exist)
	insertSampleData(db)

	// Save the data back to the JSON file
	if err := saveDataToJSON(dbPath, db); err != nil {
		log.Fatal(err)
	}

	/***
	// Query the data and print the results
	fmt.Println("ID\tWord\tWeight")
	fmt.Println("------------------")
	for _, word := range db.Words {
		fmt.Printf("%d\t%s\t%d\n", word.ID, word.Word, word.Weight)
	}***/

	testORWords(db)
	fmt.Printf("\n")
	testXORWords(db)
	fmt.Printf("\n")
	testANDWords(db)
	fmt.Printf("\n")

}

func Test(db *Database) {
	// Example usage of searchByID
	idToSearch := 2
	matchingRowsByID := searchByID(db, idToSearch)
	fmt.Println("Matching rows with ID", idToSearch, ":")
	for _, row := range matchingRowsByID {
		fmt.Println(row)
	}

	// Example usage of searchByWord
	wordToSearch := "gold"
	matchingRowsByWord := searchByWord(db, wordToSearch)
	fmt.Println("Matching rows with word", wordToSearch, ":")
	for _, row := range matchingRowsByWord {
		fmt.Println(row)
	}
}

// loadDataFromJSON reads the JSON file and loads data into the provided database.
func loadDataFromJSON(filename string, db *Database) error {
	file, err := os.Open(filename)
	if err != nil {
		// If the file does not exist, return an empty database without an error
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	return json.Unmarshal(byteValue, db)
}

// saveDataToJSON saves the database data to the JSON file.
func saveDataToJSON(filename string, db *Database) error {
	data, err := json.MarshalIndent(db, "", "    ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0644)
}

// wordExists checks if a word with the given ID already exists in the database.
func wordExists(db *Database, wordID int, wordText string) bool {
	for _, word := range db.Words {
		if (word.ID == wordID) && (word.Word == wordText) {
			return true
		}
	}
	return false
}

// insertSampleData inserts some example data into the database.
func insertSampleData(db *Database) {
	words := []Word{
		{1, "active", 1},
		{2, "ai", 1},
		{3, "alive", 1},
		{4, "all", 1},
		{5, "balance", 1},
		{6, "build", 1},
		{7, "cash", 1},
		{8, "diy", 1},
		{9, "donate", 1},
		{10, "fast", 1},
		{11, "friend", 1},
		{12, "gold", 1},
		{13, "good", 1},
		{14, "keep", 1},
		{15, "learn", 1},
		{16, "line", 1},
		{17, "low", 1},
		{18, "move", 1},
		{19, "nothing", 1},
		{20, "nuke", 1},
		{21, "private", 1},
		{22, "saving", 1},
		{23, "sell", 1},
		{24, "small", 1},
		{25, "war", 1},
		{26, "water", 1},
		{27, "patience", 1},
		{28, "wisdom", 1},
		{29, "heart", 1},
		{30, "love", 1},
		{31, "joy", 1},
		{32, "truth", 1},
		{33, "married", 1},
		{34, "mourn", 1},
		{35, "stress", 1},
		{36, "seed", 1},
		{37, "health", 1},
		{38, "bright", 1},
		{39, "eternal", 1},
		{40, "early", 1},
		{41, "care", 1},
		{42, "give", 1},
		{43, "fire", 1},
		{44, "soil", 1},
		{45, "real", 1},
		{46, "bitter", 1},
		{47, "inner", 1},
		{48, "trust", 1},
		{49, "calm", 1},
		{50, "need", 1},
		{51, "progress", 1},
		{52, "smart", 1},
		{53, "one", 1},
		{54, "defend", 1},
		{55, "product", 1},
		{56, "invest", 1},
		{57, "short", 1},
		{58, "strong", 1},
		{59, "faith", 1},
		{60, "solve", 1},
		{61, "courage", 1},
		{62, "fear", 1},
		{63, "just", 1},
		{64, "hope", 1},
		{65, "spirit", 1},
		{66, "holy", 1},

		{1, "passive", 0},
		{2, "manual", 0},
		{3, "death", 0},
		{4, "group", 0},
		{5, "bias", 0},
		{6, "destroy", 0},
		{7, "gold", 0},
		{8, "buy", 0},
		{9, "spend", 0},
		{10, "slow", 0},
		{11, "enemy", 0},
		{12, "fiat", 0},
		{13, "evil", 0},
		{14, "burn", 0},
		{15, "share", 0},
		{16, "circle", 0},
		{17, "high", 0},
		{18, "still", 0},
		{19, "everything", 0},
		{20, "gun", 0},
		{21, "public", 0},
		{22, "donate", 0},
		{23, "buy", 0},
		{24, "big", 0},
		{25, "peace", 0},
		{26, "gold", 0},
		{27, "pride", 0},
		{28, "money", 0},
		{29, "brain", 0},
		{30, "hate", 0},
		{31, "mourn", 0},
		{32, "lie", 0},
		{33, "lonely", 0},
		{34, "party", 0},
		{35, "happy", 0},
		{36, "fruit", 0},
		{37, "wealth", 0},
		{38, "dark", 0},
		{39, "temporary", 0},
		{40, "late", 0},
		{41, "iqnore", 0},
		{42, "take", 0},
		{43, "water", 0},
		{44, "metal", 0},
		{45, "fake", 0},
		{46, "sweet", 0},
		{47, "outer", 0},
		{48, "doubt", 0},
		{49, "anxious", 0},
		{50, "want", 0},
		{51, "result", 0},
		{52, "brute", 0},
		{53, "many", 0},
		{54, "attack", 0},
		{55, "consume", 0},
		{56, "trade", 0},
		{57, "long", 0},
		{58, "weak", 0},
		{59, "seen", 0},
		{60, "problem", 0},
		{61, "timid", 0},
		{62, "bold", 0},
		{63, "sacrifice", 0},
		{64, "despair", 0},
		{65, "flesh", 0},
		{66, "wicked", 0},
	}

	//for _, word := range words {
	//	db.Words = append(db.Words, word)
	//}
	for _, newWord := range words {
		if !wordExists(db, newWord.ID, newWord.Word) {
			db.Words = append(db.Words, newWord)
		}
	}
}

func testORWords(db *Database) {
	// Test cases
	word1 := "active"
	word2 := "passive"
	resid, restext, err := bitwiseORWords(word1, word2, db)
	if resid > 0 {
		fmt.Printf("%s OR %s = %s ( %d )\n", word1, word2, restext, resid)
	} else {
		fmt.Printf("%s OR %s = %s ( %d )\n", word1, word2, err, resid)
	}

	// Test cases
	word1 = "circle"
	word2 = "line"
	resid, restext, err = bitwiseORWords(word1, word2, db)
	if resid > 0 {
		fmt.Printf("%s OR %s = %s ( %d )\n", word1, word2, restext, resid)
	} else {
		fmt.Printf("%s OR %s = %s ( %d )\n", word1, word2, err, resid)
	}

	// Test cases
	word1 = "build"
	word2 = "destroy"
	resid, restext, err = bitwiseORWords(word1, word2, db)
	if resid > 0 {
		fmt.Printf("%s OR %s = %s ( %d )\n", word1, word2, restext, resid)
	} else {
		fmt.Printf("%s OR %s = %s ( %d )\n", word1, word2, err, resid)
	}

	// Test cases
	word1 = "evil"
	word2 = "good"
	resid, restext, err = bitwiseORWords(word1, word2, db)
	if resid > 0 {
		fmt.Printf("%s OR %s = %s ( %d )\n", word1, word2, restext, resid)
	} else {
		fmt.Printf("%s OR %s = %s ( %d )\n", word1, word2, err, resid)
	}

	// Test cases
	word1 = "gold"
	word2 = "fiat"
	resid, restext, err = bitwiseORWords(word1, word2, db)
	if resid > 0 {
		fmt.Printf("%s OR %s = %s ( %d )\n", word1, word2, restext, resid)
	} else {
		fmt.Printf("%s OR %s = %s ( %d )\n", word1, word2, err, resid)
	}

	// Test cases
	word1 = "gold"
	word2 = "water"
	resid, restext, err = bitwiseORWords(word1, word2, db)
	if resid > 0 {
		fmt.Printf("%s OR %s = %s ( %d )\n", word1, word2, restext, resid)
	} else {
		fmt.Printf("%s OR %s = %s ( %d )\n", word1, word2, err, resid)
	}

	// Test cases
	word1 = "mourn"
	word2 = "happy"
	resid, restext, err = bitwiseORWords(word1, word2, db)
	if resid > 0 {
		fmt.Printf("%s OR %s = %s ( %d )\n", word1, word2, restext, resid)
	} else {
		fmt.Printf("%s OR %s = %s ( %d )\n", word1, word2, err, resid)
	}
}

func testXORWords(db *Database) {
	// Test cases
	word1 := "active"
	word2 := "passive"
	resid, restext, err := bitwiseXORWords(word1, word2, db)
	if resid > 0 {
		fmt.Printf("%s XOR %s = %s ( %d )\n", word1, word2, restext, resid)
	} else {
		fmt.Printf("%s XOR %s = %s ( %d )\n", word1, word2, err, resid)
	}

	// Test cases
	word1 = "circle"
	word2 = "line"
	resid, restext, err = bitwiseXORWords(word1, word2, db)
	if resid > 0 {
		fmt.Printf("%s XOR %s = %s ( %d )\n", word1, word2, restext, resid)
	} else {
		fmt.Printf("%s XOR %s = %s ( %d )\n", word1, word2, err, resid)
	}

	// Test cases
	word1 = "build"
	word2 = "destroy"
	resid, restext, err = bitwiseXORWords(word1, word2, db)
	if resid > 0 {
		fmt.Printf("%s XOR %s = %s ( %d )\n", word1, word2, restext, resid)
	} else {
		fmt.Printf("%s XOR %s = %s ( %d )\n", word1, word2, err, resid)
	}

	// Test cases
	word1 = "evil"
	word2 = "good"
	resid, restext, err = bitwiseXORWords(word1, word2, db)
	if resid > 0 {
		fmt.Printf("%s XOR %s = %s ( %d )\n", word1, word2, restext, resid)
	} else {
		fmt.Printf("%s XOR %s = %s ( %d )\n", word1, word2, err, resid)
	}

	// Test cases
	word1 = "gold"
	word2 = "fiat"
	resid, restext, err = bitwiseXORWords(word1, word2, db)
	if resid > 0 {
		fmt.Printf("%s XOR %s = %s ( %d )\n", word1, word2, restext, resid)
	} else {
		fmt.Printf("%s XOR %s = %s ( %d )\n", word1, word2, err, resid)
	}

	// Test cases
	word1 = "gold"
	word2 = "water"
	resid, restext, err = bitwiseXORWords(word1, word2, db)
	if resid > 0 {
		fmt.Printf("%s XOR %s = %s ( %d )\n", word1, word2, restext, resid)
	} else {
		fmt.Printf("%s XOR %s = %s ( %d )\n", word1, word2, err, resid)
	}

	// Test cases
	word1 = "mourn"
	word2 = "happy"
	resid, restext, err = bitwiseXORWords(word1, word2, db)
	if resid > 0 {
		fmt.Printf("%s XOR %s = %s ( %d )\n", word1, word2, restext, resid)
	} else {
		fmt.Printf("%s XOR %s = %s ( %d )\n", word1, word2, err, resid)
	}
}

func testANDWords(db *Database) {
	// Test cases
	word1 := "active"
	word2 := "passive"
	resid, restext, err := bitwiseANDWords(word1, word2, db)
	if resid > 0 {
		fmt.Printf("%s AND %s = %s ( %d )\n", word1, word2, restext, resid)
	} else {
		fmt.Printf("%s AND %s = %s ( %d )\n", word1, word2, err, resid)
	}

	// Test cases
	word1 = "circle"
	word2 = "line"
	resid, restext, err = bitwiseANDWords(word1, word2, db)
	if resid > 0 {
		fmt.Printf("%s AND %s = %s ( %d )\n", word1, word2, restext, resid)
	} else {
		fmt.Printf("%s AND %s = %s ( %d )\n", word1, word2, err, resid)
	}

	// Test cases
	word1 = "build"
	word2 = "destroy"
	resid, restext, err = bitwiseANDWords(word1, word2, db)
	if resid > 0 {
		fmt.Printf("%s AND %s = %s ( %d )\n", word1, word2, restext, resid)
	} else {
		fmt.Printf("%s AND %s = %s ( %d )\n", word1, word2, err, resid)
	}

	// Test cases
	word1 = "evil"
	word2 = "good"
	resid, restext, err = bitwiseANDWords(word1, word2, db)
	if resid > 0 {
		fmt.Printf("%s AND %s = %s ( %d )\n", word1, word2, restext, resid)
	} else {
		fmt.Printf("%s AND %s = %s ( %d )\n", word1, word2, err, resid)
	}

	// Test cases
	word1 = "active"
	word2 = "active"
	resid, restext, err = bitwiseANDWords(word1, word2, db)
	if resid > 0 {
		fmt.Printf("%s AND %s = %s ( %d )\n", word1, word2, restext, resid)
	} else {
		fmt.Printf("%s AND %s = %s ( %d )\n", word1, word2, err, resid)
	}

	// Test cases
	word1 = "circle"
	word2 = "circle"
	resid, restext, err = bitwiseANDWords(word1, word2, db)
	if resid > 0 {
		fmt.Printf("%s AND %s = %s ( %d )\n", word1, word2, restext, resid)
	} else {
		fmt.Printf("%s AND %s = %s ( %d )\n", word1, word2, err, resid)
	}

	// Test cases
	word1 = "build"
	word2 = "build"
	resid, restext, err = bitwiseANDWords(word1, word2, db)
	if resid > 0 {
		fmt.Printf("%s AND %s = %s ( %d )\n", word1, word2, restext, resid)
	} else {
		fmt.Printf("%s AND %s = %s ( %d )\n", word1, word2, err, resid)
	}

	// Test cases
	word1 = "evil"
	word2 = "evil"
	resid, restext, err = bitwiseANDWords(word1, word2, db)
	if resid > 0 {
		fmt.Printf("%s AND %s = %s ( %d )\n", word1, word2, restext, resid)
	} else {
		fmt.Printf("%s AND %s = %s ( %d )\n", word1, word2, err, resid)
	}
}
