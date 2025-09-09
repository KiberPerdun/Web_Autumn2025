package repository

import (
	"fmt"
	"strconv"
	"strings"
)

type Repository struct {
}

func NewRepository() (*Repository, error) {
	return &Repository{}, nil
}

type Order struct {
	ID             int
	Title          string
	Description    string
	Composition    string
	ImageURL       string
	CorpusSize     string
	ServiceWords   string
	FrequencyWords []Word
}

type Word struct {
	Text      string
	Frequency float64
}

type Author struct {
	ID             int
	Name           string
	Composition    string
	Description    string
	ImageURL       string
	CorpusSize     string
	ServiceWords   string
	FrequencyWords []Word
}

func (r *Repository) GetOrders() ([]Order, error) {
	orders := []Order{
		{
			ID:           1,
			Title:        "Лев Толстой",
			Description:  "Один из наиболее известных русских писателей и мыслителей, один из величайших в мире писателей‑романистов. Участник обороны Севастополя.",
			Composition:  "романы, рассказы, письма",
			ImageURL:     "https://upload.wikimedia.org/wikipedia/commons/c/c6/L.N.Tolstoy_Prokudin-Gorsky.jpg",
			CorpusSize:   "7 Мб",
			ServiceWords: "предлоги, союзы, местоимения, артикли",
			FrequencyWords: []Word{
				{"и", 8.2},
				{"в", 6.1},
				{"он", 4.5},
				{"но", 3.0},
			},
		},
		{
			ID:           2,
			Title:        "Фёдор Достоевский",
			Description:  "Русский писатель, мыслитель, философ и публицист. Член-корреспондент Петербургской академии наук. Классик мировой литературы.",
			Composition:  "романы, повести, рассказы",
			ImageURL:     "https://upload.wikimedia.org/wikipedia/commons/thumb/7/78/Vasily_Perov_-_Портрет_Ф.М.Достоевского_-_Google_Art_Project.jpg/800px-Vasily_Perov_-_Портрет_Ф.М.Достоевского_-_Google_Art_Project.jpg",
			CorpusSize:   "6.8 Мб",
			ServiceWords: "предлоги, союзы, местоимения, частицы",
			FrequencyWords: []Word{
				{"и", 7.8},
				{"в", 5.9},
				{"что", 4.7},
				{"не", 3.5},
			},
		},
		{
			ID:           3,
			Title:        "Антон Чехов",
			Description:  "Русский писатель, прозаик, драматург, публицист, врач. Классик мировой литературы. Один из самых известных драматургов мира.",
			Composition:  "рассказы, пьесы, повести",
			ImageURL:     "https://upload.wikimedia.org/wikipedia/commons/b/bc/%D0%90._%D0%9F._%D0%A7%D0%B5%D1%85%D0%BE%D0%B2%2C_%D0%AF%D0%BB%D1%82%D0%B0.jpg",
			CorpusSize:   "5.2 Мб",
			ServiceWords: "предлоги, союзы, частицы, местоимения",
			FrequencyWords: []Word{
				{"и", 7.2},
				{"в", 5.6},
				{"на", 4.1},
				{"с", 3.2},
			},
		},
	}

	if len(orders) == 0 {
		return nil, fmt.Errorf("массив пустой")
	}

	return orders, nil
}

func (r *Repository) GetOrderByID(id string) (*Order, error) {
	orders, err := r.GetOrders()
	if err != nil {
		return nil, err
	}

	orderID, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("неверный формат ID: %v", err)
	}

	for _, order := range orders {
		if order.ID == orderID {
			return &order, nil
		}
	}

	return nil, fmt.Errorf("заказ с ID %s не найден", id)
}

func (r *Repository) GetAuthors() ([]Author, error) {
	authors := []Author{
		{
			ID:           1,
			Name:         "Лев Толстой",
			Composition:  "романы, рассказы, письма",
			Description:  "Один из наиболее известных русских писателей и мыслителей, один из величайших в мире писателей-романистов. Участник обороны Севастополя.",
			ImageURL:     "https://upload.wikimedia.org/wikipedia/commons/c/c6/L.N.Tolstoy_Prokudin-Gorsky.jpg",
			CorpusSize:   "9.3 Мб",
			ServiceWords: "предлоги, союзы, частицы, местоимения",
			FrequencyWords: []Word{
				{"и", 7.9},
				{"в", 5.8},
				{"что", 4.2},
				{"он", 3.7},
			},
		},
		{
			ID:           2,
			Name:         "Фёдор Достоевский",
			Composition:  "романы, рассказы, письма",
			Description:  "Русский писатель, мыслитель, философ и публицист. Член-корреспондент Петербургской академии наук с 1877 года. Классик мировой литературы, по данным ЮНЕСКО, один из самых читаемых писателей в мире.",
			ImageURL:     "https://upload.wikimedia.org/wikipedia/commons/thumb/7/78/Vasily_Perov_-_Портрет_Ф.М.Достоевского_-_Google_Art_Project.jpg/800px-Vasily_Perov_-_Портрет_Ф.М.Достоевского_-_Google_Art_Project.jpg",
			CorpusSize:   "8.2 Мб",
			ServiceWords: "предлоги, союзы, междометия, местоимения",
			FrequencyWords: []Word{
				{"и", 8.2},
				{"в", 6.1},
				{"не", 5.3},
				{"что", 4.1},
			},
		},
		{
			ID:           3,
			Name:         "Антон Чехов",
			Composition:  "Рассказ, пьесы, письма",
			Description:  "Русский писатель, прозаик, драматург, публицист, врач, общественный деятель в сфере благотворительности. Классик мировой литературы. Почётный академик Императорской академии наук по разряду изящной словесности. Один из самых известных драматургов мира",
			ImageURL:     "https://upload.wikimedia.org/wikipedia/commons/b/bc/%D0%90._%D0%9F._%D0%A7%D0%B5%D1%85%D0%BE%D0%B2%2C_%D0%AF%D0%BB%D1%82%D0%B0.jpg",
			CorpusSize:   "5.7 Мб",
			ServiceWords: "предлоги, союзы, частицы, артикли",
			FrequencyWords: []Word{
				{"и", 6.8},
				{"в", 5.4},
				{"на", 3.9},
				{"с", 3.1},
			},
		},
	}

	if len(authors) == 0 {
		return nil, fmt.Errorf("массив авторов пуст")
	}

	return authors, nil
}

func (r *Repository) GetAuthorByID(id string) (*Author, error) {
	authors, err := r.GetAuthors()
	if err != nil {
		return nil, err
	}

	authorID, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("неверный формат ID: %v", err)
	}

	for _, author := range authors {
		if author.ID == authorID {
			return &author, nil
		}
	}

	return nil, fmt.Errorf("автор с ID %s не найден", id)
}

func (r *Repository) SearchAuthors(query string) ([]Author, error) {
	authors, err := r.GetAuthors()
	if err != nil {
		return nil, err
	}

	if query == "" {
		return authors, nil
	}

	query = strings.ToLower(query)
	var filteredAuthors []Author

	for _, author := range authors {
		if strings.Contains(strings.ToLower(author.Name), query) {
			filteredAuthors = append(filteredAuthors, author)
			continue
		}

		if strings.Contains(strings.ToLower(author.Composition), query) {
			filteredAuthors = append(filteredAuthors, author)
			continue
		}

		if strings.Contains(strings.ToLower(author.Description), query) {
			filteredAuthors = append(filteredAuthors, author)
			continue
		}
	}

	return filteredAuthors, nil
}
