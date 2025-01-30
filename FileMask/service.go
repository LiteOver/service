package FileMask

type Service struct {
	Prod Producer
	Pres Presenter
}

func NewService(prod Producer, pres Presenter) *Service {
	return &Service{Prod: prod, Pres: pres}
}

func (s *Service) Mask(data []byte) string {
	var prefix string = "http://"
	for i := range data {
		if i < len(data)-len(prefix) && string(data[i:i+len(prefix)]) == prefix {
			for i < len(data)-len(prefix) && (data[i+len(prefix)] != 32) {
				data[i+len(prefix)] = 42
				i++
			}
		}
	}
	return string(data)
}
