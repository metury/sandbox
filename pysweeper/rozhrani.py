import model
import tkinter
import hra

class GrafickeRozhrani:
    """vykreslování pole a práce s grafickou stránkou"""
    def __init__(self, vyska, sirka, b):
        self.sirka = sirka
        self.vyska = vyska
        self.pocet_bomb = b
        self.viditelne_pole = [["?" for i in range(self.sirka)] for j in range(self.vyska)]

    def set_viditelne_pole(self, x, y, znak):
        """na danou pozici se dá ekvivalent z pole bomb a čísel"""
        self.viditelne_pole[x][y] = znak

    def set_akce(self, akce):
        """tady není potřeba, ale pro opuštění podmínky se uloží
           a používá v advanced grafickému rozhraní"""
        self.akce = akce

    def get_znak_z_viditelneho(self, x, y):
        if x >= 0 and x < self.vyska and y >= 0 and y < self.sirka:
            return self.viditelne_pole[x][y]
        else:
            return None
  
    def set_kompletni_pole(self, pole):
        """pro zviditelnění celého pole"""
        self.viditelne_pole = pole

    def tiskni_vyhru(self):
        print("\n","-"*12,"\n","|Vyhrál si!|\n","-"*12,"\n")

    def tiskni_prohru(self):
        print("\n","-"*5,"\n","|BUM|\n","-"*5,"\n")
        
    def vykresli(self, pocet_vlajek = 0):
        """co se vykreslí po každé akci včetně počet bomb a volných vlajek"""
        print(" "*3, end = "")
        for i in range(self.sirka):
            print(chr(ord("A")+i), end=" ")
        print()
        for i in range(self.vyska):
            print(" "*2, end ="")
            print(" -"*self.sirka)
            print("%2d" %(i+1),end="")
            print("|",end = "")
            print("|".join(self.viditelne_pole[i]), end = "")
            print("|")
        print(f"Počet bomb: {self.pocet_bomb}")
        print(f"Počet volných vlajek: {pocet_vlajek}")

class GrafickeRozhrani1(GrafickeRozhrani):
    """nové prostředí pomocí tkinteru, generace tlačítek
       vše se mění tady až na akce v imaginárním poli,
       kde jsou uložené data bomb, to jde přes akci"""
    
    def tiskni_prohru(self):
        self.text3.config(text = "BUM")
        self.newgabt.grid(row = self.vyska+1, column = 4)
        self.vlajka.config(bg = "grey")
        self.odkryt.config(bg = "grey")
        self.pole.title("!!PROHRÁL SI!!")

    def tiskni_vyhru(self):
        self.text3.config(text = "VÝHRA")
        self.newgabt.grid(row = self.vyska+1, column = 4)
        self.vlajka.config(bg = "grey")
        self.odkryt.config(bg = "grey")
        self.pole.title("!!VYHRÁL SI!!")

    def vykresli(self,pocetvlajek):
        """slovo se volá do akce, která je napojená
           a popřípadě se mění, přes tlačítka"""
        self.slovo = "O"

        def nova_hra():
            """po zmáčknutí tlačítka pro start
               nové hry"""
            self.pole.destroy()
            h0 = hra.Hra()
            h0.set_grafika("A")
            h0.start()

        def set_slovo(slovo):
            """změna slova aby se provedla správná akce,
               popř. konec okna"""
            self.slovo = slovo
            if self.slovo == "V":
                self.vlajka.config(bg = "yellow")
                self.odkryt.config(bg = "grey")
            elif self.slovo == "O":
                self.vlajka.config(bg = "grey")
                self.odkryt.config(bg = "yellow")
            else:
                self.akce.obecna_akce(self.slovo,0,0)
                self.pole.destroy()


        def klik(i,j):
            """provedení dané akce a vykreslení správných znaků"""
            if self.akce.get_konec() == False:
                self.akce.obecna_akce(self.slovo, i, j)
                self.pocet_vlajek.config(text = self.akce.get_pocet_vlajek())
                colors = ["black","blue","green","red","dark blue", "dark green", "orange", "black", "dark red"]
                for y,elem in enumerate(self.btn):
                    for x,button in enumerate(elem):
                        if str(self.viditelne_pole[y][x]) == "0":
                            color = colors[0]
                            button.config(text = " ", fg = color)
                            continue
                        elif str(self.viditelne_pole[y][x]) == "B" or str(self.viditelne_pole[y][x]) == "?" \
                            or str(self.viditelne_pole[y][x]) == "V":
                            color = "black"
                        else:
                            color = colors[int(self.viditelne_pole[y][x])]
                        button.config(text = str(self.viditelne_pole[y][x]), fg = color)

        self.btn = []
        self.pole = tkinter.Tk()
        self.pole.title("Minesweeper")

        prazdny_policko = tkinter.Label(text = " ", width = 5)
        prazdny_policko.grid(row = 0, column = 0)

        #vytvoření elementů v okně
        for k in range(self.sirka):
            pismeno = tkinter.Label(text = chr(ord("A")+k), width = 5)
            pismeno.grid(row = 0, column = k+1)
        for i in range(self.vyska):
            cislo = tkinter.Label(text = i+1, width = 5)
            cislo.grid(row = i+1, column = 0)
            temp = []
            for j in range(self.sirka):
                bt = tkinter.Button(text = "?", fg = "black", width = 5, command = lambda c = i, d = j: klik(c,d))
                temp.append(bt)
                bt.grid(row = i+1, column = j+1)
            self.btn.append(temp)
  
        self.vlajka = tkinter.Button(text = "Vlajka", command = lambda: set_slovo("V"), bg = "grey")
        self.odkryt = tkinter.Button(text = "Odkrýt", command = lambda: set_slovo("O"), bg = "yellow")
        self.konec = tkinter.Button(text = "Konec", command = lambda: set_slovo("K"), bg = "red")
        self.newgabt = tkinter.Button(text = "START", bg = "green", command = lambda: nova_hra())

        self.text1 = tkinter.Label(text = "Vlajky")
        self.text2 = tkinter.Label(text = "Bomby")
        self.text3 = tkinter.Label(fg = "red")
        self.pocet_vlajek = tkinter.Label(text = self.akce.get_pocet_vlajek())
        self.pocet_bomb = tkinter.Label(text = self.pocet_bomb)

        self.vlajka.grid(row = self.vyska + 1, column = 1)
        self.odkryt.grid(row = self.vyska + 1, column = 2)
        self.konec.grid(row = self.vyska + 1, column = 3)
        self.text1.grid(row = self.vyska +1, column = 6)
        self.pocet_vlajek.grid(row = self.vyska +1, column = 7)
        self.text2.grid(row = self.vyska +1, column = 8)
        self.pocet_bomb.grid(row = self.vyska +1, column = 9)
        self.text3.grid(row = self.vyska +1, column = 5)

        self.pole.mainloop()

class Vlastnosti:
    """ve vlastnostech se ze vstupu zadají data a předají do hry/pole
    včetně kontroly vstupu, aby byli hodnoty platné, popřípadě se změní na bližší hranici""" 
    def __init__(self):
        self.vyska_horni_mez = 30
        self.vyska_dolni_mez = 3
        self.sirka_horni_mez = 26
        self.sirka_dolni_mez = 3
        self.minimum_bomb = 1
        self.maximum_bomb = 3

    def kontrolax(self, x):
        """kontrola a popř změna výšky"""
        temp_x1 = min(x, self.vyska_horni_mez)
        temp_x2 = max(x, self.vyska_dolni_mez)
        if x != temp_x1 or x != temp_x2:
            if x < temp_x2:
                x = temp_x2
            else:
                x = temp_x1
            
            print(f"Musel jsem změnit výšku na {x}.")
        return x

    def kontrolay(self, y):
        """-||- šířka"""
        temp_y1 = min(y, self.sirka_horni_mez)
        temp_y2 = max(y, self.sirka_dolni_mez)
        if y != temp_y1 or y != temp_y2:
            if y < temp_y2:
                y = temp_y2
            else:
                y = temp_y1
            print(f"Musel jsem změnit šířku na {y}.")
        return y

    def kontrolab(self, b):
        """bomb"""
        if b > self.maximum_bomb or b < self.minimum_bomb:
            b = self.maximum_bomb//2
            print(f"Změnil jsem počet bomb na průměr max/min {b}.")
        return b

    def vstup(self):
        try:
            x = int(input(f"Zadej výšku pole (rozmezí od {self.vyska_dolni_mez} do {self.vyska_horni_mez}): "))
            x = self.kontrolax(x)
        except ValueError:
            print("Zadej číslo.")
            return 0
    
        try:
            y = int(input(f"Zadej šířku pole (rozmezí od {self.sirka_dolni_mez} do {self.sirka_horni_mez}): "))
            y = self.kontrolay(y)
        except ValueError:
            print("Zadej číslo.")
            return 0

        self.maximum_bomb = x*y

        try:
            b = int(input(f"Zadej počet bomb (rozmezí od {self.minimum_bomb} do {self.maximum_bomb}): "))
            b = self.kontrolab(b)
        except ValueError:
            print("Zadej číslo.")
            return 0

        Pole0 = model.Pole(x,y,b)
        return Pole0

class Vlastnosti1(Vlastnosti):
    """Pozměněný vstup pomocí grafické vizualizace skrz knihovnu tkinter"""
    def vstup(self):
        def uloz_vstup(event):
            """použií vstupu od uživatele"""
            try:
                self.x = int(vyskatx.get())
                self.y = int(sirkatx.get())
                self.b = int(bombytx.get())
                self.x = self.kontrolax(self.x)
                self.y = self.kontrolay(self.y)
                self.maximum_bomb = self.x*self.y
                self.b = self.kontrolab(self.b)
                okno.destroy()
            except ValueError:
                print("Zadej čísla!")
        
        def rychla_hra(event):
            """přednastavené udáje na novou hru"""
            self.x = 20
            self.y = 20
            self.b = 40
            okno.destroy()

        """vytvoření okna pomocí tkinteru"""       
        okno = tkinter.Tk()
        okno.title("Hodnoty")
        #vytvoření elementů
        vyskalb = tkinter.Label(okno, width = 25, text = "Výška pole (3 - 30):")
        sirkalb = tkinter.Label(okno, width = 25, text = "Šířka pole (3 - 26):")
        bombylb = tkinter.Label(okno, width = 25, text = "Počet bomb (1 - obsah pole):")
        vyskatx = tkinter.Entry(width = 15)
        sirkatx = tkinter.Entry(width = 15)
        bombytx = tkinter.Entry(width = 15)
        startbt = tkinter.Button(okno, text = "START", width = 15)
        faststart = tkinter.Button(okno, text = "RYCHLÁ HRA (20|20|40)", width = 25)
        #přidání na dané okno
        vyskalb.grid(row = 0, column = 0)
        vyskatx.grid(row = 0, column = 1)
        sirkalb.grid(row = 1, column = 0)
        sirkatx.grid(row = 1, column = 1)
        bombylb.grid(row = 2, column = 0)
        bombytx.grid(row = 2, column = 1)
        startbt.grid(row = 3, column = 1)
        faststart.grid(row = 3, column = 0)
        #navázání funkcí
        faststart.bind("<Button-1>", rychla_hra)
        startbt.bind("<Button-1>", uloz_vstup)

        okno.mainloop()

        pole0 = model.Pole(self.x, self.y, self.b, False)
        return pole0
