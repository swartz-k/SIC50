import { Nav } from '@douyinfe/semi-ui'
import useTranslation from 'hooks/useTranslation'
import { useNavigate } from 'react-router'

export default () => {
  const nav = useNavigate()
  const [t] = useTranslation()

  const menuItems = [
    {
      itemKey: '#/overview',
      text: t('overview'),
      // icon: <OverviewIcon />,
      onClick: () => nav('/overview'),
    },
    {
      itemKey: '#/task',
      text: t('task'),
      // icon: <TaskIcon />,
      onClick: () => nav('/task'),
    },
    {
      itemKey: '#/result',
      text: t('quer_result'),
      // icon: <TaskIcon />,
      onClick: () => nav('/result'),
    },
  ]

  return (
    <Nav
      defaultSelectedKeys={[window.location.hash]}
      style={{ maxWidth: 150, height: '100%' }}
      items={menuItems}
      footer={{
        collapseText: () => t('collapse'),
        collapseButton: true,
      }}
    />
  )
}
