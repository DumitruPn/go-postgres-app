package notification

type Notification struct {
	ID    int
	Value string
}

type CreateNotificationRequest struct {
	Value string `json:"value"`
}

type Dto struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

func AsDto(notification Notification) Dto {
	return Dto{
		ID:    notification.ID,
		Value: notification.Value,
	}
}

func AsDtos(notifications []Notification) []Dto {
	dtos := []Dto{}
	for _, notification := range notifications {
		dtos = append(dtos, AsDto(notification))
	}
	return dtos
}
