function getTrValues(idTr){

    var tr = document.getElementById(idTr);
    var cells =  tr.cells
    var cellcount = cells.length;
    var cell = [];
    j=0;

    for( c=0; c<cellcount; c++) {
        cell[j]= cells[c].innerHTML;
        console.log(cell[j]);
        j++;
    }

    $.ajax({
        url: 'receive',
        type: 'post',
        dataType: 'html',
        data : { code: cell[0], libelle: cell[1], quantite: cell[2], prix:  cell[3], date: cell[4], categorie: cell[5], seuil: cell[6]},
        success : function(data) {
            $('#result').html(data);
        },
    });
    document.location.href="Détailsduproduit";
}

function OperatRetour(idTr){
    var tr = document.getElementById(idTr);
    var cells =  tr.cells
    var cellcount = cells.length;
    var cell = [];
    var descript= document.getElementById("description").value;
    j=0;

    for( c=0; c<cellcount; c++) {
        cell[j]= cells[c].innerHTML;
        j++;
    }
    $.ajax({
        url: 'receiveSortInst',
        type: 'post',
        dataType: 'html',
        data :{ code: cell[0], libelle: cell[1], quantite: cell[2], prix:  cell[3], date: cell[4], categorie: cell[5], benef : cell[6], idOP: cell[8], description: descript},
        success : function(data) {
            $('#result').html(data);
        },
    });
    tr.remove();
}