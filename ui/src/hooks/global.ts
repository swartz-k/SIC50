import { createGlobalState } from 'react-hooks-global-state'

const initialStats = {
  task: [],
  taskLoading: false,
}

const { useGlobalState } = createGlobalState(initialStats)
export default useGlobalState
