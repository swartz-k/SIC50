export interface ITask {
  CreatedAt: string
  ID: number
  UpdatedAt: string
  // config: {input_layer: "serving_default_input_input", output_layer: "StatefulPartitionedCall",…}
  input_layer: string
  output: string
  output_layer: string
  // steps: {0: {concentration: 0.666666,…}, 1: {concentration: 0,…}, 2: {concentration: 0.111111, images: null}}
  status: string
  task_id: string
}

export interface ITaskResponse {
  task_id: string
}
