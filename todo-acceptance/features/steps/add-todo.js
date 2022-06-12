const {Given, When, Then} = require("cucumber");
const openUrl = require("../support/action/openUrl");
const waitForSelector = require("../support/action/waitForSelector");
const assert = require("assert");

Given(/^Empty ToDo list$/, async function () {
    await openUrl.call(this, "")
});

When(/^I write "([^"]*)" to text box and click to add button$/, async function(inputText) {
    const inputSelector = "#todo-input"
    await waitForSelector.call(this, inputSelector, 5)
    await this.page.type(inputSelector, inputText)
    await this.page.click("#todo-add-button")
});

Then(/^I should see "([^"]*)" item in ToDo list$/, async function (todoText) {
    this.page.on('console', msg => {
        for (let i = 0; i < msg.args().length; ++i)
            console.log(`${i}: ${msg.args()[i]}`);
    });
    const todoItemSelector = "#todo-text"
    const todoItem = await this.page.$$eval(
        todoItemSelector,
        async (elements, todoText) => {
            return elements.find(el => el.textContent === todoText)
        },
        todoText
    )
    console.log(todoItem)
    assert.equal(!!todoItem, true)
});
