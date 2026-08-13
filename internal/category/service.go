package category

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetAll() ([]CategoryResponse, error) {
	categories, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (s *Service) GetByID(id uint) (*Category, error) {
	category, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	if category == nil {
		return nil, ErrCategoryNotFound
	}

	return category, nil
}

func (s *Service) Create(category *Category) (*Category, error) {
	createdCategory, err := s.repository.Create(*category)
	if err != nil {
		return nil, err
	}

	return createdCategory, nil
}

func (s *Service) Update(id uint, data *Category) (*Category, error) {
	category, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	if category == nil {
		return nil, ErrCategoryNotFound
	}

	category.Name = data.Name
	category.Slug = data.Slug
	category.Description = data.Description
	category.IsActive = data.IsActive

	updatedCategory, err := s.repository.Update(category)
	if err != nil {
		return nil, err
	}

	return updatedCategory, nil
}

func (s *Service) Delete(id uint) error {
	category, err := s.repository.FindByID(id)
	if err != nil {
		return err
	}

	if category == nil {
		return ErrCategoryNotFound
	}

	return s.repository.DeleteByID(id)
}