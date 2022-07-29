import { HashRouter, Routes, Route } from 'react-router-dom'

import Overview from 'pages/Overview'
import Tasks from 'pages/Task'
import CommonLayout from 'components/CommonLayout'
import Result from 'pages/Result'

const Router = () => {
  return (
    <HashRouter>
      <CommonLayout>
        <Routes>
          <Route path='/' element={<Tasks />} />
          <Route path='/overview' element={<Overview />} />
          <Route path='/result' element={<Result />} />
          <Route path='*' element={<Tasks />} />
        </Routes>
      </CommonLayout>
    </HashRouter>
  )
}

export default Router
