import model
import rozhrani

class Akce:
    """hráč může buď dát vlajku nebo odhalit pole a také ukončit hru,
       je tady i pocet bomb na kterých nejsou vlajky a počet volných vlajek"""
    def __init__(self, grafika, pole, bomb):
        self.konec = False
        self.grafika = grafika
        self.pole = pole
        self.pocet_bomb = bomb
        self.pocet_vlajek = bomb
        
    def obecna_akce(self, slovo, x, y):
        """pro volání jakékoliv akce, kontorluje se vstup až tady"""
        if slovo == "V":
            self.vlajka(x,y)
        elif slovo == "O":
            self.klik(x,y)
        elif slovo == "K":
            self.konec = True
            self.grafika.set_kompletni_pole(self.pole.get_pole())
        
        if self.pocet_bomb < 1:
            self.konec = True
            self.grafika.set_kompletni_pole(self.pole.get_pole())
            self.grafika.tiskni_vyhru()

    def get_pocet_vlajek(self):
        return self.pocet_vlajek

    def get_konec(self):
        return self.konec

    def vlajka(self, x, y):
        """přidání vlajky a pokud se treí na pole, kde je bomba,
           tak se odebere počet bomb a hlídá se počet vlajek, 
           taky lze jen na neznáme pole"""
        if self.pocet_vlajek > 0 and self.grafika.get_znak_z_viditelneho(x,y) == "?":
            self.grafika.set_viditelne_pole(x,y,"V")
            self.pocet_vlajek -= 1
            if (self.pole.get_znak(x,y)=="B"):
                self.pocet_bomb -= 1
        else:
            print("Všechny vlajky použity.")

    def klik(self, x, y):
        """po kliknutí se odkryje co je na poli a pokud bomba tak je konec,
           pokud nula, tak otevře všechny nuly okolo včetně krajních čísel"""
        if self.pole.get_znak(x,y) != None:

            if self.grafika.get_znak_z_viditelneho(x,y) == "V":
                self.pocet_vlajek = self.pocet_vlajek + 1

            self.grafika.set_viditelne_pole(x,y,str(self.pole.get_znak(x,y)))

            if self.pole.get_znak(x,y) == "B":
                self.konec = True
                self.grafika.tiskni_prohru()
                self.grafika.set_kompletni_pole(self.pole.get_pole())

            elif self.pole.get_znak(x,y) == 0:
                okoli_pole = [(x+1, y+1),(x+1, y),(x, y+1),(x+1, y-1),(x-1, y+1),(x, y-1),(x-1, y),(x-1, y-1)]
                for prvek in okoli_pole:
                    if self.grafika.get_znak_z_viditelneho(prvek[0],prvek[1]) == "?":
                        self.klik(prvek[0], prvek[1])

class Hra:
    """samotné spuštění hry a nekonečná iterace pro akce 
       s možností hrát novou hru, kontrola vstupu pro akce
       změny pokud se jedná o terminál nebo advanced,
       pokude je advanced, tak opakování je na grafice""" 

    def set_grafika(self, st):
        self.terminal = st
        if self.terminal == "T":
            self.vlast = rozhrani.Vlastnosti()
        else:
            self.vlast = rozhrani.Vlastnosti1()

    def start(self):
        pole1 = self.vlast.vstup()
        while pole1 == 0:
            pole1 = self.vlast.vstup()
        grafika1 = pole1.grafika()
        akce1 = Akce(grafika1, pole1, pole1.get_pocet_bomb())
        #pokud je advanced, tak je třeba propojit s akci
        grafika1.set_akce(akce1)
        grafika1.vykresli(akce1.get_pocet_vlajek())

        if self.terminal == "T":
            while akce1.get_konec() == False:
                slovo = input("Zadej akci (V - vlajka, O - odkrýt, K - konec): ")
                if slovo == "V" or slovo == "O":
                    souradnice = input("Zadej souřadnice (písmeno a číslo oddělené mezerou): ").split()
                    try:
                        y = ord((souradnice[0]))-ord("A")
                    except TypeError:
                        print("Chyba v zápisu.")
                        continue
                    try:    
                        x = int(souradnice[1])-1
                    except:
                        print("Chyba v zápisu.")
                        continue
                else:
                    x = 0
                    y = 0

                if slovo == "K" or slovo == "V" or slovo == "O":
                    akce1.obecna_akce(slovo,x,y)
                    grafika1.vykresli(akce1.get_pocet_vlajek())
            
            restart = input("Nová hra? (A / N): ")
            if restart == "A":
                self.start()
            while restart != "A" and restart != "N":
                print("Buď A = ano, nebo N = ne.")
                restart = input("Nová hra? (A / N): ")
                if restart == "A":
                    self.start()
