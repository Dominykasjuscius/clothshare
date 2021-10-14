<template>
    <div>
        <b-card-group deck v-for="(row,index) in productMatrix" :key="index" align="center">
            <b-card v-for="column in row" 
                :key="column._id" 
                class="mb-3" 
                border-variant="primary" 
                img-alt="image"
                style="min-width: 25%;"
                :img-src="column.image" 
                img-top header-bg-variant="primary"
                v-visible="!column.invisible"
               >
                <template #header>
                    <h4 class="mb-0">{{column._id}}</h4>
                </template>

                <b-card-body class="body">
                    <b-card-title>{{column.price}}€</b-card-title>
                    <b-card-sub-title class="mb-2">{{column.size}}</b-card-sub-title>
                    <b-card-text>iosdadsiodimodsmoidsimosdfmiofdsmoidfmoifsomifsoimsfmifsmoifsoimfsoimsef</b-card-text>
                </b-card-body>
                <b-card-body>
                    <b-link :to="{ name: 'Product', params:{id: column._id}}" class="card-link">View product
                    </b-link>
                    <router-view ></router-view>
                    <a class="card-link">Another link </a>
                </b-card-body>
            </b-card>
        </b-card-group>
    </div>
</template>

<script>
export default {
    props:{
        productMatrix: [],
    }
}
</script>

<style lang="scss">
.card-link {
    color: var(--blue);
}
</style>