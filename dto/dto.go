package dto

type Health struct {
  clusterName string `json:"cluster_name"`
  status string `json:"status"`                                                                                                                                                                                        
  status bool `json:"timed_out"`                                                                                                                                                                                        
  status uint `json:"number_of_nodes"`                                                                                                                                                                                      
  status string `json:"number_of_data_nodes": 1,                                                                                                                                                                                 
  status string `json:"active_primary_shards": 0,                                                                                                                                                                                
  status string `json:"active_shards": 0,                                                                                                                                                                                        
  status string `json:"relocating_shards": 0,                                                                                                                                                                                    
  status string `json:"initializing_shards": 0,                                                                                                                                                                                  
  status string `json:"unassigned_shards": 0,                                                                                                                                                                                    
  status string `json:"delayed_unassigned_shards": 0,                                                                                                                                                                            
  status string `json:"number_of_pending_tasks": 0,                                                                                                                                                                              
  status string `json:"number_of_in_flight_fetch": 0,                                                                                                                                                                            
  status string `json:"task_max_waiting_in_queue_millis": 0,                                                                                                                                                                     
  status string `json:"active_shards_percent_as_number": 100 
}