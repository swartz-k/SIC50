import { Button, Dropdown, Layout, Nav, Tag, Typography } from '@douyinfe/semi-ui'
import { IconTranslate } from 'assets/imgs/IconTranslate'
import useTranslation from 'hooks/useTranslation'
import i18n from 'i18n'
import { useNavigate } from 'react-router'

export default () => {
  const [t] = useTranslation()
  const { Header } = Layout
  const { Text } = Typography
  const nav = useNavigate()
  // const username = Cookies.get('username')

  return (
    <Header>
      <Nav mode='horizontal' style={{ height: '40px' }}>
        <Nav.Header>
          <Text link onClick={() => nav('/')}>
            SIC50
          </Text>
        </Nav.Header>
        <Nav.Footer>
          <Text link={{ target: '_blank', href: 'paperpaper' }} style={{ marginRight: '15px' }}>
            {t('docs')}
          </Text>
          <Dropdown
            trigger={'hover'}
            position={'bottomRight'}
            render={
              <Dropdown.Menu>
                <Dropdown.Item>
                  <Button theme='borderless' onClick={() => i18n.changeLanguage('zh')}>
                    中文
                  </Button>
                </Dropdown.Item>
                <Dropdown.Item>
                  <Button theme='borderless' onClick={() => i18n.changeLanguage('en')}>
                    English
                  </Button>
                </Dropdown.Item>
              </Dropdown.Menu>
            }
          >
            <Tag style={{ background: 'transparent' }}>
              <IconTranslate />
            </Tag>
          </Dropdown>
        </Nav.Footer>
      </Nav>
    </Header>
  )
}
