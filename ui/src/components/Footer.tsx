import { Layout } from '@douyinfe/semi-ui'

export default () => {
  const { Footer } = Layout
  return (
    <Footer
      style={{
        height: '2%',
        display: 'flex',
        justifyContent: 'space-between',
        color: 'var(--semi-color-text-2)',
        padding: '20px',
      }}
    >
      <div style={{ margin: '0 auto' }}>
        <span>Copyright 2022 SIC50. </span>
      </div>
    </Footer>
  )
}
