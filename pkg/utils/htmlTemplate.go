package utils

var HTMLTemplate = `<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>验证码</title>
  <style>
    body { margin: 0; padding: 0; font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Helvetica, Arial, sans-serif; background-color: #f6f8fa; color: #24292f; }
    table { border-spacing: 0; width: 100%%; max-width: 600px; margin: 0 auto; }
    td { padding: 0; }
    @media screen and (max-width: 600px) {
      .container { width: 100%% !important; padding: 15px !important; }
    }
  </style>
</head>
<body style="margin: 0; padding: 40px 0; background-color: #f6f8fa;">
  <table role="presentation" cellspacing="0" cellpadding="0" border="0" align="center" style="background-color: #ffffff; border-radius: 16px; overflow: hidden; box-shadow: 0 4px 24px rgba(0,0,0,0.05); border: 1px solid #eaeaea;">
    <tr>
      <td style="height: 6px; background: #333333;"></td>
    </tr>

    <tr>
      <td style="padding: 40px 40px 10px 40px; text-align: center;">
        <h1 style="margin: 0; font-size: 22px; font-weight: 700; color: #1f2328; letter-spacing: -0.5px;">%s</h1>
        <p style="margin: 10px 0 0; font-size: 15px; color: #57606a;">请验证您的身份以继续.</p>
      </td>
    </tr>

    <tr>
      <td style="padding: 20px 40px;">
        <div style="background-color: #f6f8fa; border-radius: 12px; padding: 24px; text-align: center; border: 1px solid #edf2f7;">
          <p style="margin: 0 0 10px; font-size: 12px; text-transform: uppercase; color: #8c959f; font-weight: 600;">验证码</p>
          
          <div style="font-family: 'SF Mono', 'Segoe UI Mono', 'Courier New', monospace; font-size: 32px; font-weight: 700; letter-spacing: 6px; color: #24292f; margin-bottom: 8px;">
            %s
          </div>
          
          <p style="margin: 5px 0 0; font-size: 13px; color: #cf222e; font-weight: 500;">验证码5分钟后过期</p>
        </div>
      </td>
    </tr>

    <tr>
      <td style="padding: 10px 40px 40px 40px;">
        <table width="100%%" style="border-top: 1px solid #eaeaea; padding-top: 20px;">
          <tr>
            <td style="font-size: 13px; color: #6e7781;">您的邮件地址:</td>
            <td style="font-size: 13px; color: #24292f; font-weight: 500; text-align: right;">%s</td>
          </tr>
          <tr>
            <td style="font-size: 13px; color: #6e7781; padding-top: 8px;">发送时间:</td>
            <td style="font-size: 13px; color: #24292f; font-weight: 500; text-align: right; padding-top: 8px;">%s</td>
          </tr>
        </table>
      </td>
    </tr>

  </table>
  
  <div style="text-align: center; margin-top: 24px; font-size: 12px; color: #8c959f;">
    <p style="margin: 0;">&copy; %d %s. All rights reserved.</p>
  </div>

</body>
</html>`

//const MARKDOWN_CONTENT = `
//## 第一章节：起源与思考
//
//在这个章节中，我们深入探讨技术的本质。随着 **Vue 3.0** 的发布，组合式 API 带来了逻辑复用的新范式。
//
//### 核心概念
//
//1.  **响应式系统**：基于 Proxy 的高性能实现
//2.  **组合式 API**：更灵活的代码组织方式
//3.  **Teleport**：内置的模态框解决方案
//
//## 第二章节：构建现代化体验
//
//设计不仅仅是视觉的堆砌。Awwwards 风格强调的是一种流动的体验（Flow）。
//
//> "好的设计是显而易见的，伟大的设计是透明的。"
//
//## 代码示例
//
//\`\`\`javascript
//function hello() {
//console.log("Hello World");
//}
//\`\`\`
//
//## 第三章节：未来的展望
//
//这不仅是一个博客，更是一个技术游乐场。我们期待看到更多创新的交互方式。
//`
