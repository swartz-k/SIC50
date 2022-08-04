import { Button, Card, Col, Descriptions, Input, Layout, Row, Tag } from '@douyinfe/semi-ui'
import useTranslation from 'hooks/useTranslation'
import { IconSearch, IconClose, IconRefresh } from '@douyinfe/semi-icons'
import { useState } from 'react'
import axios from 'axios'
import { useCallback } from 'react'
import { useSearchParams } from 'react-router-dom'
import { useTasks } from 'hooks/useTask'
import { ITask } from 'types/ITask'
import { LineChart, Line, CartesianGrid, XAxis, YAxis, Tooltip } from 'recharts'

export default () => {
  const [t] = useTranslation()
  const { tasks, removeTask } = useTasks()
  const [searchParams, setSearchParams] = useSearchParams()

  const [searchLoading, setSearchLoading] = useState(false)
  const [task, setTask] = useState<ITask>()

  const queryTaskId = useCallback(async () => {
    const id = searchParams.get('id')
    setSearchLoading(true)
    try {
      const resp = await axios.get<ITask>('/api/v1/task', {
        params: {
          id: id,
        },
      })
      setTask(resp.data)
    } finally {
      setSearchLoading(false)
    }
  }, [searchParams])

  return (
    <Layout>
      <Card style={{ margin: '20px' }}>
        <Row>
          <Col span={16} offset={4}>
            <Input
              placeholder='Enter your task id'
              value={searchParams.get('id') ?? ''}
              onChange={(v) => setSearchParams({ id: v })}
            />
          </Col>
          <Col span={4}>
            <Button
              theme='borderless'
              icon={<IconSearch />}
              loading={searchLoading}
              style={{ width: 120, marginLeft: 4 }}
              onClick={() => queryTaskId()}
            >
              {t('search')}
            </Button>
          </Col>
        </Row>
        {!task ? (
          <Row>
            <Col span={22} offset={1} style={{ marginTop: '20px' }}>
              {tasks().map((tid) => {
                if (tid.length > 0) {
                  return (
                    <Tag
                      onClick={() => setSearchParams({ id: tid })}
                      key={tid}
                      color='blue'
                      type='light'
                      style={{ marginRight: '5px', marginBottom: '5px' }}
                      closable
                      onClose={() => removeTask(tid)}
                    >
                      {tid}
                    </Tag>
                  )
                }
              })}
            </Col>
          </Row>
        ) : (
          <div style={{ marginTop: '29px' }}>
            <Card
              title={t('result')}
              headerExtraContent={
                <div>
                  <Button theme='borderless' loading={searchLoading} onClick={queryTaskId}>
                    <IconRefresh />
                  </Button>
                  <Button theme='borderless' onClick={() => setTask(undefined)}>
                    <IconClose />
                  </Button>
                </div>
              }
            >
              <Descriptions>
                <Descriptions.Item itemKey='ID'>{task.ID}</Descriptions.Item>
                <Descriptions.Item itemKey={t('create_at')}>{task.CreatedAt}</Descriptions.Item>
                <Descriptions.Item itemKey={t('status')}>{task.status}</Descriptions.Item>
              </Descriptions>

              {task.status === 'SUCCESS' && (
                <LineChart
                  width={600}
                  height={300}
                  data={task.config.output.result}
                  margin={{ top: 5, right: 20, bottom: 5, left: 0 }}
                >
                  <Line type='monotone' dataKey='res' stroke='#8884d8' />
                  <CartesianGrid stroke='#ccc' strokeDasharray='5 5' />
                  <XAxis dataKey='con' />
                  <YAxis />
                  <Tooltip />
                </LineChart>
              )}
            </Card>
          </div>
        )}
      </Card>
    </Layout>
  )
}
