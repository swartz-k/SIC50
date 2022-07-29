/* eslint-disable @typescript-eslint/no-unused-vars */
import { Button, Card, Col, Input, InputGroup, Layout, Row, Select, Tag, Typography } from '@douyinfe/semi-ui'
import useTranslation from 'hooks/useTranslation'
import { IconSearch } from '@douyinfe/semi-icons'
import { useState } from 'react'
import axios from 'axios'
import { useCallback } from 'react'
import { useSearchParams } from 'react-router-dom'
import { useTasks } from 'hooks/useTask'
import { ITask } from 'types/ITask'

export default () => {
  const { Text } = Typography
  const [t] = useTranslation()
  const { tasks } = useTasks()
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
                    >
                      {tid}
                    </Tag>
                  )
                }
              })}
            </Col>
          </Row>
        ) : (
          <Text>{task.CreatedAt}</Text>
        )}
      </Card>
    </Layout>
  )
}
