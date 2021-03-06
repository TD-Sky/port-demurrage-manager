import dayjs from "dayjs"

/** 格式化时间 */
export const formatDateTime = (time: any) => {
  if (time == null || time === "") {
    return "N/A"
  }
  const date = new Date(time)
  return dayjs(date).format("YYYY-MM-DD HH:mm:ss")
}

/** 将全局 css 导入 js 中使用 */
export const getCssVariableValue = (cssVariableName: string) => {
  let cssVariableValue = ""
  try {
    // 没有拿到值时，会返回空串
    cssVariableValue = getComputedStyle(document.documentElement).getPropertyValue(cssVariableName)
  } catch (error) {
    console.error(error)
  }
  return cssVariableValue
}

// 断言未来的日期
export function isnt_future(time: Date): boolean {
    return time.getTime() <= Date.now()
}

// 断言开始计费
export function fees_incurred(duration: number) {
    return duration > 0
}
