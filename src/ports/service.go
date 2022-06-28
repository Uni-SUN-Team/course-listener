package ports

import "unisun/api/course-listener/src/model"

type ConsumerService interface {
	GetInformationFormStrapi(payloadRequest model.ServiceIncomeRequest) (string, error)
	GetAdvisorInfomation(ids string) (string, error)
}
