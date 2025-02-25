import Cookies from 'js-cookie'
import { useCallback } from 'react'

export const useTasks = () => {
  const tasks = useCallback(() => {
    const _tasks = Cookies.get('task_ids') ?? ''
    return _tasks.split(',')
  }, [])

  const addTask = useCallback((task_id: string) => {
    const _tasks = Cookies.get('task_ids') ?? ''
    const tasks = _tasks.split(',')
    Cookies.set('task_ids', [task_id, ...tasks].join(','))
  }, [])

  const removeTask = useCallback((task_id: string) => {
    const _tasks = Cookies.get('task_ids') ?? ''
    const tasks = _tasks.split(',')
    return tasks.map((tid) => tid !== task_id)
  }, [])

  return {
    tasks,
    addTask,
    removeTask,
  }
}
