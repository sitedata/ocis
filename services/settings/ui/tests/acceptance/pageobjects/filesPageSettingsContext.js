module.exports = {
  commands: {
    getMenuList: async function () {
      const menu = []
      await this.isVisible('@openNavigationBtn', (res) => {
        if (res.value) {
          this.click('@openNavigationBtn')
        }
      })
      await this.waitForElementVisible('@fileSidebarNavItem')
      await this.api
        .elements('@fileSidebarNavItem', result => {
          result.value.map(item => {
            this.api.elementIdText(item.ELEMENT, res => {
              menu.push(res.value)
            })
            return undefined
          })
        })
      return menu
    },
    getUserMenu: async function () {
      const menu = []
      await this
        .waitForElementVisible('@userMenuBtn')
        .click('@userMenuBtn')
        .waitForElementVisible('@userMenuContainer')
      await this.api
        .elements('@userMenuItem', result => {
          result.value.map(item => {
            this.api.elementIdText(item.ELEMENT, res => {
              menu.push(res.value)
            })
            return undefined
          })
        })
      await this
        .click('@userMenuBtn')
        .waitForElementNotPresent('@userMenuContainer')
      return menu
    },
    getFileHeaderItems: async function () {
      const menu = []
      await this.waitForElementVisible('@fileTableHeaderItems')
      await this.api
        .elements('@fileTableHeaderItems', result => {
          result.value.map(item => {
            this.api.elementIdText(item.ELEMENT, res => {
              menu.push(res.value)
            })
            return undefined
          })
        })
      return menu
    }
  },

  elements: {
    pageHeader: {
      selector: '.oc-page-title'
    },
    fileSidebarNavItem: {
      selector: '.oc-sidebar-nav-item'
    },
    openNavigationBtn: {
      selector: '.oc-app-navigation-toggle'
    },
    userMenuBtn: {
      selector: '#_userMenuButton'
    },
    userMenuItem: {
      selector: '#account-info-container li'
    },
    userMenuContainer: {
      selector: '#account-info-container'
    },
    fileTableHeaderItems: {
      selector: '//*[@id="files-space-table"]//th[not(.//div)]',
      locateStrategy: 'xpath'
    }
  }
}
