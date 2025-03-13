import { After, Before, Given, Then, When } from '@cucumber/cucumber'
import type { BrowserContext } from '@playwright/test'
import { type Browser, type Page, chromium, expect } from '@playwright/test'

const adjectives = [
  'happy',
  'sad',
  'bright',
  'dark',
  'fast',
  'slow',
  'strong',
  'weak',
  'loud',
  'quiet',
  'big',
  'small',
  'soft',
  'hard',
  'warm',
  'cold',
  'sharp',
  'dull',
  'brave',
  'shy',
  'rich',
  'poor',
  'funny',
  'serious',
  'kind',
  'rude',
  'friendly',
  'mean',
  'smart',
  'dumb',
  'fresh',
  'stale',
  'clean',
  'dirty',
  'safe',
  'dangerous',
  'light',
  'heavy',
  'simple',
  'complex',
  'beautiful',
  'ugly',
  'new',
  'old',
  'deep',
  'shallow',
  'high',
  'low',
  'strong',
  'fragile',
]

const nouns = [
  'dog',
  'cat',
  'house',
  'car',
  'tree',
  'river',
  'mountain',
  'ocean',
  'cloud',
  'star',
  'book',
  'pen',
  'phone',
  'table',
  'chair',
  'window',
  'door',
  'floor',
  'wall',
  'ceiling',
  'computer',
  'laptop',
  'keyboard',
  'mouse',
  'screen',
  'cup',
  'plate',
  'spoon',
  'fork',
  'knife',
  'bag',
  'box',
  'shirt',
  'pants',
  'shoes',
  'hat',
  'watch',
  'clock',
  'bottle',
  'glass',
  'road',
  'bridge',
  'island',
  'city',
  'village',
  'farm',
  'forest',
  'desert',
  'beach',
  'hill',
]

function randomUsername() {
  return `${pick(adjectives)}_${pick(nouns)}_${new Date().getTime()}@example.com`
}

function randomPassword() {
  return `${pick(nouns)}_${pick(nouns)}_${pick(nouns)}`
}

function pick<T>(data: T[]) {
  return data[Math.floor(Math.random() * data.length)]
}

Before(async function () {
  this.browser = await chromium.launch({ headless: false })

  this.context = await this.browser.newContext()

  this.page = await this.context.newPage()
})

After(async function () {
  const page = this.page as Page
  const browserContent = this.context as BrowserContext
  const browser = this.browser as Browser

  await page.close()
  await browserContent.close()
  await browser.close()
})

Given('a registered user with a valid email and password', async function () {
  this.username = 'registered@example.com'
  this.password = 'does-not-matter'

  try {
    const response = await fetch('http://localhost:8080/v1/auth/register', {
      body: JSON.stringify({
        email: this.username,
        password: this.password,
      }),
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
    })

    if (!response.ok && response.status != 403) {
      throw response.body
    }
  } catch (e) {
    console.error('fail to create scenario', e)
    throw e
  }
})

When('they enter their credentials correctly', async function () {
  const page = this.page as Page

  await page.goto('http://localhost:3000/login')
  await page.getByRole('textbox', { name: 'อีเมล' }).click()
  await page.getByRole('textbox', { name: 'อีเมล' }).fill(this.username)
  await page.getByRole('textbox', { name: 'อีเมล' }).press('Tab')
  await page.getByRole('textbox', { name: 'รหัสผ่าน', exact: true }).click()
  await page
    .getByRole('textbox', { name: 'รหัสผ่าน', exact: true })
    .fill(this.password)
  await page
    .getByRole('textbox', { name: 'รหัสผ่าน', exact: true })
    .press('Tab')
  await page.getByRole('button', { name: 'เข้าสู่ระบบ', exact: true }).click()

  await page.waitForResponse('http://localhost:8080/v1/auth/login')
})

Then('the system logs them in', async function () {
  const page = this.page as Page

  await page.goto('http://localhost:3000/')

  await page.waitForResponse('http://localhost:8080/v1/auth/me')
  const locators = await page
    .getByRole('button', { name: 'สวัสดี, Registered' })
    .all()

  expect(locators.length).toBe(1)
})

When('they enter their credentials incorrectly', async function () {
  const page = this.page as Page

  await page.goto('http://localhost:3000/login')
  await page.getByRole('textbox', { name: 'อีเมล' }).fill(this.username)
  await page
    .getByRole('textbox', { name: 'รหัสผ่าน', exact: true })
    .fill('wrong-password')
  await page.getByRole('button', { name: 'เข้าสู่ระบบ', exact: true }).click()

  await page.waitForResponse('http://localhost:8080/v1/auth/login')
})

Then('the system does not log them in', async function () {
  const page = this.page as Page
  const locators = await page.getByText('Incorrect email or password').all()

  expect(locators.length).toBe(1)
})

Given('a user with a unique email', async function () {
  this.username = randomUsername()
  this.password = randomPassword()
})

When('they submit their email and password', async function () {
  const page = this.page as Page

  await page.goto('http://localhost:3000/')
  await page.getByRole('link', { name: 'สมัครใช้งาน' }).click()
  await page.getByRole('textbox', { name: 'ชื่อจริง' }).fill('John')
  await page.getByRole('textbox', { name: 'นามสกุล' }).fill('Doe')
  await page
    .getByRole('textbox', { name: 'เบอร์โทรศัพท์' })
    .fill('086-666-6666')
  await page.getByRole('combobox', { name: 'เพศ' }).click()
  await page.getByRole('option', { name: 'หญิง' }).click()
  await page.getByRole('button', { name: 'ถัดไป' }).click()
  await page.getByRole('textbox', { name: 'อีเมล' }).fill(this.username)
  await page
    .getByRole('textbox', { name: 'รหัสผ่าน', exact: true })
    .fill(this.password)
  await page
    .getByRole('textbox', { name: 'ยืนยันรหัสผ่าน' })
    .fill(this.password)
  await page.getByRole('button', { name: 'สมัครใช้งาน', exact: true }).click()

  await page.waitForResponse('http://localhost:8080/v1/auth/register')
})

Then('the system creates their account successfully', async function () {
  const page = this.page as Page

  await page.goto('http://localhost:3000/')

  await page.waitForResponse('http://localhost:8080/v1/auth/me')
  const locators = await page
    .getByRole('button', { name: 'สวัสดี, John' })
    .all()

  expect(locators.length).toBe(1)
})

Given('a user with an already registered email', async function () {
  this.username = 'this_is_duplicate_email@example.com'
  this.password = 'does-not-matter'

  try {
    const response = await fetch('http://localhost:8080/v1/auth/register', {
      body: JSON.stringify({
        email: this.username,
        password: this.password,
      }),
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
    })

    if (!response.ok && response.status != 403) {
      throw response.body
    }
  } catch (e) {
    console.error('fail to create scenario', e)
    throw e
  }
})

When('they attempt to register with the same email', async function () {
  const page = this.page as Page

  await page.goto('http://localhost:3000/')
  await page.getByRole('link', { name: 'สมัครใช้งาน' }).click()
  await page.getByRole('textbox', { name: 'ชื่อจริง' }).fill('John')
  await page.getByRole('textbox', { name: 'นามสกุล' }).fill('Doe')
  await page
    .getByRole('textbox', { name: 'เบอร์โทรศัพท์' })
    .fill('086-666-6666')
  await page.getByRole('combobox', { name: 'เพศ' }).click()
  await page.getByRole('option', { name: 'หญิง' }).click()
  await page.getByRole('button', { name: 'ถัดไป' }).click()
  await page.getByRole('textbox', { name: 'อีเมล' }).fill(this.username)
  await page
    .getByRole('textbox', { name: 'รหัสผ่าน', exact: true })
    .fill(this.password)
  await page
    .getByRole('textbox', { name: 'ยืนยันรหัสผ่าน' })
    .fill(this.password)
  await page.getByRole('button', { name: 'สมัครใช้งาน', exact: true }).click()

  await page.waitForResponse('http://localhost:8080/v1/auth/register')
})

Then('the system shows an error message', async function () {
  const page = this.page as Page
  const locators = await page.getByText('Something went wrong').all()

  expect(locators.length).toBe(1)
})
