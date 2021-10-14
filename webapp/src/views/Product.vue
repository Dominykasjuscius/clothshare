<template>
    <b-container class="mt-3">
        <b-row>
            <b-col md="8">
                <b-card title="Card Title" :img-src="product.image" img-alt="Image" img-top tag="article" class="mb-3">
                    <b-card-text>
                        {{product._id}}
                        Some quick example text to build on the card title and make up the bulk of the card's
                        content.
                    </b-card-text>
                    <b-button href="#" variant="secondary">Go somewhere</b-button>
                </b-card>
            </b-col>
            <b-col md="4">
                <description-card :product="product">
                </description-card>
            </b-col>
        </b-row>
        <b-row >
            <b-col md="8">
          <product-deck :productMatrix="productRows">
                </product-deck>
            </b-col>
            <b-col sm="4">
                <user-card :userid="product.author">
                </user-card>
            </b-col>
        </b-row>
        <b-row>
            <b-col md="8">
                <pagination :totalRows="userProducts.length" align-h="center"
                    :currentPage="currentPage" :pageSize="pageSize" v-on:page:update="updatePage"></pagination>
            </b-col>
        </b-row>
    </b-container>
</template>

<script lang="ts">
    import axios from 'axios'
    import Vue from 'vue';
    import DescriptionCard from '../components/DescriptionCard.vue'
    import User from '../components/UserCard.vue'
    import ProductDeck from '../components/ProductDeck.vue'
    import Pagination from '../components/Pagination.vue'
    export default Vue.extend({
        data() {
            return {
                id: "",
                image: "",
                product: Object as any,
                userProducts:[] as any,
                userVisibleProducts:[] as any,
                productRows: [] as any,
                currentPage: 1,
                pageSize: 4,              
                rowSize: 3,  
            }
        },
        components: {
            "description-card": DescriptionCard,
            "user-card": User,
            "product-deck": ProductDeck,
            "pagination": Pagination
        },
        methods: {
            async fetchProduct(id:string) {
                axios.get("/api/products/" + id)
                    .then(function (this:any, response:any) {
                        this.product = response.data
                    }.bind(this))
                    .then(data => this.populateProduct())
                    .then(data => this.fetchUserProducts())
                    .then(data => this.product.createdAt = this.formatDate(this.product.createdAt))
            },

            async fetchUserProducts() {
                axios.get("/api/products/user/" + this.product.author)
                    .then(function (this :any, response:any) {
                        this.userProducts = response.data
                    }.bind(this))
                    .then(data => this.populateProduct())
                    .then(data => this.updateVisibleProducts())
                    .catch(err => console.log(err.message))
            },
            updatePage(page: number) {
                this.currentPage = page
                this.updateVisibleProducts()
            },

            updateVisibleProducts() {
                this.userProducts = this.userProducts.filter((obj: { _id: any; }) => obj._id !== this.product._id);
                this.userVisibleProducts = this.userProducts.slice((this.currentPage -1) * this.pageSize, this.pageSize * (this.currentPage))
                
                this.fillEmptySpaces()

                this.productRows = []
                try {
                  while(this.userVisibleProducts.length) this.productRows.push(this.userVisibleProducts.splice(0,this.rowSize));
                }                
                catch (error){
                    console.log(error)
                }
            },

            fillEmptySpaces() {
                if (this.userVisibleProducts.length % this.rowSize != 0) {
                    let productsToAddCount = Math.abs(this.userVisibleProducts.length -  this.pageSize)
                    let prodLength = this.userVisibleProducts.length
                    for (let i = 0; i < productsToAddCount; i++) {
                        var invisibleProd = Object.assign({}, this.userVisibleProducts[0])

                        invisibleProd.invisible = true
                        this.userVisibleProducts[prodLength + i] = invisibleProd
                    }   
                }  
            },
            populateProduct() {
                for (var key of Object.keys(this.product)) {
                    if (this.product[key] == "") {
                        this.product[key] = "unknown"
                    }
                }
            },
            formatDate(d:string) {
                var date = new Date(d)
                var dd = date.getDate()
                var mm = date.getMonth() + 1;
                var yyyy = date.getFullYear();
                if (dd < 10) {
                    dd = Number('0') + Number(dd) 
                }
                if (mm < 10) {
                    mm = Number('0') + Number(mm) 
                }
                return d = yyyy + "-" + mm + "-" + dd
            }
        },

        watch: {
            '$route.params.id': {
                handler(id) {
                    const pid= id
                    this.fetchProduct(pid)
                },
                immediate: true,
            }
        },
        mounted() {
            this.fetchProduct(this.$route.params.id)
        }

        
    })
</script>

<style>
    .sidecol {
        font-size: 1.3rem !important;
        vertical-align: middle;
    }
    .card-deck .card {
    max-width: calc(33.3% - 30px);
    }

    .description-vars {
        font-weight: bold;
    }
    .row-eq-height {
  display: -webkit-box;
  display: -webkit-flex;
  display: -ms-flexbox;
  display:         flex;
}
</style>