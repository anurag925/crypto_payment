package app

import "github.com/anurag925/crypto_payment/utils/task"

func WorkerObj() *task.Client {
	return instance.Worker().Instance()
}
