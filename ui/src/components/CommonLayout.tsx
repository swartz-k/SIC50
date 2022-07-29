import React from 'react'
import { Layout } from '@douyinfe/semi-ui'
import Header from './Header'
import Sidebar from './Sidebar'
import Footer from './Footer'

export interface ICommonLayoutProps {
  children: React.ReactNode
}

export default ({ children }: ICommonLayoutProps) => {
  const { Content } = Layout

  return (
    <Layout style={{ height: '100vh', backgroundColor: 'var(--semi-color-bg-1)' }}>
      <Layout.Header>
        <Header />
      </Layout.Header>
      <Layout>
        <Layout.Sider>
          <Sidebar />
        </Layout.Sider>
        <Layout style={{ backgroundColor: 'var(--semi-color-bg-1)' }}>
          <Layout.Content>
            <Content>{children}</Content>
            <Layout.Footer>
              <Footer />
            </Layout.Footer>
          </Layout.Content>
        </Layout>
      </Layout>
    </Layout>
  )
}
