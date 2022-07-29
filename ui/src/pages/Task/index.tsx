import { Card, Layout } from '@douyinfe/semi-ui'
import CreateTaskForm from 'components/CreateTaskForm'

export default () => {
  return (
    <Layout>
      <Card style={{ margin: '20px' }}>
        <CreateTaskForm />
      </Card>
    </Layout>
  )
}
