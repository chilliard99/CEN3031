describe('Test show in sim', () => {

    beforeEach(() => {
        cy.visit('/');
    });

    it('shows all the hands', () => {
        cy.get("#mat-tab-label-0-1").click();
        cy.get("#highCard").click();
        cy.get("#mat-tab-label-0-1").click();
        cy.get("#twoPair").click();
        cy.get("#mat-tab-label-0-1").click();
        cy.get("#threeOfAKind").click();
        cy.get("#mat-tab-label-0-1").click();
        cy.get("#straight").click();
        cy.get("#mat-tab-label-0-1").click();
        cy.get("#flush").click();
        cy.get("#mat-tab-label-0-1").click();
        cy.get("#fullHouse").click();
        cy.get("#mat-tab-label-0-1").click();
        cy.get("#fourOfAKind").click();
        cy.get("#mat-tab-label-0-1").click();
        cy.get("#straightFlush").click();
        cy.get("#mat-tab-label-0-1").click();
        cy.get("#royalFlush").click();
    });

  })