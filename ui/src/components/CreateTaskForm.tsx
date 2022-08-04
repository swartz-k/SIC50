/* eslint-disable @typescript-eslint/no-unused-vars */
import React, { useCallback, useState } from 'react'
import { Form, Col, Row, Button, Slider, InputNumber, Upload, ArrayField, Toast } from '@douyinfe/semi-ui'
import { IconPlusCircle, IconMinusCircle, IconUpload } from '@douyinfe/semi-icons'

import useTranslation from 'hooks/useTranslation'
import { useTasks } from 'hooks/useTask'
import axios from 'axios'
import { ITaskResponse } from 'types/ITask'

export default () => {
  const [t] = useTranslation()
  const { addTask } = useTasks()
  const [step, setStep] = useState(3)
  const [loading, setLoading] = useState(false)

  const getArray = useCallback(() => {
    return new Array(step).fill('')
  }, [step])

  const submit = useCallback(
    async (vals: Record<string, any>) => {
      setLoading(true)
      try {
        const resp = await axios.post<ITaskResponse>('/api/v1/task/async', vals)
        addTask(resp.data.task_id)
        Toast.success(`${t('task_submit_success_tooltip')} ${resp.data.task_id}`)
      } finally {
        setLoading(false)
      }
    },
    [addTask],
  )

  return (
    <Form labelPosition={'left'} style={{ padding: 10, width: '100%' }} onSubmit={(vals) => submit(vals)}>
      {/* <Row>
        <Col span={1}>
          <Label style={{ paddingTop: '12px', paddingRight: '0' }}>{t('step_num')}</Label>
        </Col>
        <Col span={8}>
          <Slider min={1} max={16} step={1} value={step} onChange={(value) => setStep(value as number)} />
        </Col>
        <Col span={4}>
          <InputNumber min={1} max={16} onChange={(v) => setStep(v as number)} style={{ width: 100 }} value={step} />
        </Col>
      </Row> */}
      {step && (
        <ArrayField field='step' initValue={getArray()}>
          {({ add, arrayFields }) => (
            <React.Fragment>
              <Button onClick={add} icon={<IconPlusCircle />} theme='light'>
                {t('add_step')}
              </Button>
              {arrayFields.map(({ key, field, remove }, i) => (
                <div key={key} style={{ width: '100%' }}>
                  <Row>
                    <Col span={12}>
                      <Form.InputNumber
                        field={`${field}[concentration]`}
                        label={`${field} ${t('step_concentration')}`}
                        style={{ width: 120, marginRight: 16 }}
                      ></Form.InputNumber>
                    </Col>
                    <Col span={4}>
                      <Button
                        type='danger'
                        theme='borderless'
                        icon={<IconMinusCircle />}
                        onClick={remove}
                        style={{ margin: 12 }}
                      ></Button>
                    </Col>
                    <Col span={12}>
                      <Form.Upload
                        multiple
                        draggable={true}
                        action='/api/v1/upload'
                        field={`${field}[upload]`}
                        label={t('step_images')}
                        dragMainText={t('click_drag_upload_text')}
                        style={{ margin: 0, padding: 0 }}
                      />
                    </Col>
                  </Row>
                </div>
              ))}
            </React.Fragment>
          )}
        </ArrayField>
      )}
      <Button type='primary' htmlType='submit' className='btn-margin-right'>
        {t('submit')}
      </Button>
    </Form>
  )
}
