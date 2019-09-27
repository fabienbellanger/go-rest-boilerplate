$(function()
{
    // Initialisation des datatables
    // -----------------------------
    $('#errorLogsTable').DataTable({
        'lengthMenu': [[100, 200, 500], [100, 200, 500]],
    });

    $('#sqlLogsTable').DataTable({
        'lengthMenu': [[100, 200, 500], [100, 200, 500]],
    });

    $('#accessLogsTable').DataTable({
        'lengthMenu': [[100, 200, 500], [100, 200, 500]],
    });

    // Affichage du dernier onglet ouvert
    // ----------------------------------
    let currentTab = window.sessionStorage.getItem('currentTab');
    if (currentTab == null)
    {
        currentTab = '#errors';
    }
    $('#logsTab a[href="' + currentTab + '"]').tab('show');

    // Sélection d'un onglet
    // ---------------------
    $('#logsTab a').on('click', function (e) {
        e.preventDefault();
        $(this).tab('show');

        // Sauvegarde de l'onglet actif
        // ----------------------------
        window.sessionStorage.setItem('currentTab', $(this).attr('href'));
    });
});
