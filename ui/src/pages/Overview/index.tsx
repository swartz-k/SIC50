import { Card, Descriptions, Layout } from '@douyinfe/semi-ui'
import axios from 'axios'
import useTranslation from 'hooks/useTranslation'
import { useCallback, useEffect, useState } from 'react'
import { IOverview } from 'types/IOverview'

export default () => {
  const [t] = useTranslation()

  const [loading, setLoading] = useState(false)
  const [overview, setOverview] = useState<IOverview>()
  const fetchOverview = useCallback(async () => {
    setLoading(true)
    try {
      const resp = await axios.get<IOverview>('/api/v1/overview')
      setOverview(resp.data)
    } finally {
      setLoading(false)
    }
  }, [])

  useEffect(() => {
    fetchOverview()
  }, [fetchOverview])

  return (
    <Layout>
      <Card style={{ margin: '20px' }} loading={loading}>
        <Descriptions row size='large'>
          <Descriptions.Item itemKey={t('total_tasks')}>{overview ? overview.tasks : 99} </Descriptions.Item>
          <Descriptions.Item itemKey={t('total_images')}>{overview ? overview.images : 99} </Descriptions.Item>
        </Descriptions>
      </Card>
    </Layout>
  )
}
