describe('Add Card', () => {

    beforeEach(() => {
        cy.visit('/');
    });

    it('Click Suits', () => {
        cy.get("#defaultcard0").click()
        cy.get("#diamond").click()
        cy.get("#ace").click()
        cy.get("#simulationLayer").click()
        cy.get("#removeAll").click()
    });

  })