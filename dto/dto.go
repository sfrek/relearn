package dto

// Health Elasticsearch cluster status response
type Health struct {
	ClusterName             string `json:"cluster_name"`
	Status                  string `json:"status"`
	TimeOut                 bool   `json:"timed_out"`
	Nodes                   uint   `json:"number_of_nodes"`
	DataNodes               uint   `json:"number_of_data_nodes"`
	ActivePrimaryShards     uint   `json:"active_primary_shards"`
	ActiveShards            uint   `json:"active_shards"`
	RelocatingShards        uint   `json:"relocating_shards"`
	IntializingShards       uint   `json:"initializing_shards"`
	UnassignedShards        uint   `json:"unassigned_shards"`
	DelayedUnassignedShards uint   `json:"delayed_unassigned_shards"`
	PendigTasks             uint   `json:"number_of_pending_tasks"`
	InFlightFetch           uint   `json:"number_of_in_flight_fetch"`
	Tmwiqm                  uint   `json:"task_max_waiting_in_queue_millis"`
	Aspan                   uint   `json:"active_shards_percent_as_number"`
}
