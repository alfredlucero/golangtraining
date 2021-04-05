package maps

type Dictionary map[string]string

// Constant errors with our own DictionaryErr type which implements the error interface.
const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word becasue it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	// Map lookups can return 2 values; second value is a boolean which indicates if key was found successfully.
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// When passing a map to a function/method, you are copying just the pointer part, not the underlying data structure that contains the data.
// Don't initialize map as empty nil value; use map[string]string{} or make(map[string]string) instead.
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	// Maps have a built-in delete function which returns nothing.
	delete(d, word)
}
