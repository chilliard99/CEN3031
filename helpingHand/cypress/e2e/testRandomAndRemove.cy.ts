describe('Random and Remove', () => {

    beforeEach(() => {
        cy.visit('/');
    });

    it('Tests Random', () => {
        cy.get("#topButtons").click()
        cy.get("#random").click()
        cy.get("#random").click()
        cy.get("#random").click()
        cy.get("#removeAll").click()
    });

  })