import random
import rozhrani

class Pole:
    """zde se vytváří hrací pole a taky 
    je zde počet bomb a údaje znaku v poli"""
    def __init__(self, x, y, b, terminal = True):
        self.vyska_pole = x
        self.sirka_pole = y
        self.pocet_bomb = b
        #vytvoreni pole s bombami a čísly
        self.pole = [["X" for i in range(self.sirka_pole)] for j in range(self.vyska_pole)]
        self.pridej_bomby()
        self.pridej_cisla()
        #bollean jaké je graficke rozhrani
        self.terminal = terminal

    def get_pocet_bomb(self):
        return self.pocet_bomb

    def get_pole(self):
        """vrácení identického pole, akorát jsou místo čísel stringy, 
           aby se dali tisknout"""
        nove_pole =[]
        for item in self.pole:
            temp = []
            for elem in item:
                temp.append(str(elem))
            nove_pole.append(temp)
        return nove_pole

    def get_znak(self, x, y):
        if x >= 0 and x < self.vyska_pole and y >= 0 and y < self.sirka_pole:
            return self.pole[x][y]
        else:
            return None


    def grafika(self):
        """vytvoření grafického prostředí pro dané pole, vrací se objekt grafika,
           generace podle toho, jestli je terminal nebo advanced"""
        if self.terminal:
            grafika0 = rozhrani.GrafickeRozhrani(self.vyska_pole, self.sirka_pole, self.pocet_bomb)
        else:
            grafika0 = rozhrani.GrafickeRozhrani1(self.vyska_pole, self.sirka_pole, self.pocet_bomb)
        return grafika0
    
    def pridej_bomby(self):
        """pomocí random dá bomby na místa kde už nejsou"""
        temp_bomb = self.pocet_bomb
        while temp_bomb > 0:
            x = random.randint(0, self.vyska_pole-1)
            y = random.randint(0, self.sirka_pole-1)
            if self.pole[x][y] == "X":
                self.pole[x][y] = "B"
                temp_bomb -= 1
    
    def pridej_cisla(self):
        """do prázdných polí se dá správné číslo jako počet bomb okolo"""
        pocet_volnych_poli = (self.vyska_pole * self.sirka_pole) - self.pocet_bomb
        i = 0
        j = 0
        
        while pocet_volnych_poli > 0:
            cislo_bomb = 0
            if  self.pole[i][j] == "X":
                #kontorola jestli nejsou bomby na vedlejších polích + jestli tam vůbec pole je
                if i > 0 and self.pole[i-1][j] == "B": #pod
                    cislo_bomb += 1

                if i < (self.vyska_pole - 1) and self.pole[i+1][j] == "B": #nad
                    cislo_bomb += 1

                if j > 0 and self.pole[i][j-1] == "B": #vlevo
                    cislo_bomb += 1

                if j < (self.sirka_pole - 1) and self.pole[i][j+1] == "B": #vpravo
                    cislo_bomb += 1

                #kontrola jestli jsou bomby na diagonalach a jestli tam je pole
                if i > 0 and j > 0 and self.pole[i-1][j-1] == "B": #vlevo nad
                    cislo_bomb += 1

                if  j< (self.sirka_pole - 1) and i < (self.vyska_pole - 1)  and self.pole[i+1][j+1] == "B": #vpravo pod
                    cislo_bomb += 1

                if i < (self.vyska_pole - 1)  and j > 0 and self.pole[i+1][j-1] == "B": #vlevo pod
                    cislo_bomb += 1

                if i > 0 and j < (self.sirka_pole - 1) and self.pole[i-1][j+1] == "B": # vpravo nad
                    cislo_bomb += 1
                
                self.pole[i][j] = cislo_bomb
                pocet_volnych_poli -= 1
        
            if j == self.sirka_pole-1:
                j = 0
                i+=1
            else:
                j+=1
