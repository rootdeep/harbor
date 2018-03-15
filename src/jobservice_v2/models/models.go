// Copyright 2018 The Harbor Authors. All rights reserved.

package models

//Parameters for job execution.
type Parameters map[string]interface{}

//JobRequest is the request of launching a job.
type JobRequest struct {
	Job *JobData `json:"job"`
}

//JobData keeps the basic info.
type JobData struct {
	Name       string       `json:"name"`
	Parameters Parameters   `json:"parameters"`
	Metadata   *JobMetadata `json:"metadata"`
}

//JobMetadata stores the metadata of job.
type JobMetadata struct {
	JobKind       string `json:"kind"`
	ScheduleDelay uint64 `json:"schedule_delay,omitempty"`
	Cron          string `json:"cron_spec,omitempty"`
	IsUnique      bool   `json:"unique"`
}

//JobStats keeps the result of job launching.
type JobStats struct {
	Stats *JobStatData `json:"job"`
}

//JobStatData keeps the stats of job
type JobStatData struct {
	JobID       string `json:"id"`
	Status      string `json:"status"`
	JobName     string `json:"name"`
	JobKind     string `json:"kind"`
	IsUnique    bool   `json:"unique"`
	RefLink     string `json:"ref_link,omitempty"`
	EnqueueTime int64  `json:"enqueue_time"`
	UpdateTime  int64  `json:"update_time"`
	RunAt       int64  `json:"run_at,omitempty"`
	CheckIn     string `json:"check_in,omitempty"`
	CheckInAt   int64  `json:"check_in_at,omitempty"`
}

//JobPoolStats represent the healthy and status of the job service.
type JobPoolStats struct {
	WorkerPoolID string   `json:"worker_pool_id"`
	StartedAt    int64    `json:"started_at"`
	HeartbeatAt  int64    `json:"heartbeat_at"`
	JobNames     []string `json:"job_names"`
	Concurrency  uint     `json:"concurrency"`
	Status       string   `json:"status"`
}