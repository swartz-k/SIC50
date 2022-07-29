import React from 'react'
import { Layout } from '@douyinfe/semi-ui'

export interface ICommonLayoutProps {
  children: React.ReactNode
}

export default ({ children }: ICommonLayoutProps) => {
  return <Layout style={{ height: '100vh', backgroundColor: 'var(--semi-color-bg-1)' }}>{children}</Layout>
}
