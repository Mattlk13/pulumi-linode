// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.linode.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Double;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetDatabasePostgresqlConfigSharedBuffersPercentage {
    private String description;
    private Double example;
    private Double maximum;
    private Double minimum;
    private Boolean requiresRestart;
    private String type;

    private GetDatabasePostgresqlConfigSharedBuffersPercentage() {}
    public String description() {
        return this.description;
    }
    public Double example() {
        return this.example;
    }
    public Double maximum() {
        return this.maximum;
    }
    public Double minimum() {
        return this.minimum;
    }
    public Boolean requiresRestart() {
        return this.requiresRestart;
    }
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDatabasePostgresqlConfigSharedBuffersPercentage defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String description;
        private Double example;
        private Double maximum;
        private Double minimum;
        private Boolean requiresRestart;
        private String type;
        public Builder() {}
        public Builder(GetDatabasePostgresqlConfigSharedBuffersPercentage defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.example = defaults.example;
    	      this.maximum = defaults.maximum;
    	      this.minimum = defaults.minimum;
    	      this.requiresRestart = defaults.requiresRestart;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetDatabasePostgresqlConfigSharedBuffersPercentage", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder example(Double example) {
            if (example == null) {
              throw new MissingRequiredPropertyException("GetDatabasePostgresqlConfigSharedBuffersPercentage", "example");
            }
            this.example = example;
            return this;
        }
        @CustomType.Setter
        public Builder maximum(Double maximum) {
            if (maximum == null) {
              throw new MissingRequiredPropertyException("GetDatabasePostgresqlConfigSharedBuffersPercentage", "maximum");
            }
            this.maximum = maximum;
            return this;
        }
        @CustomType.Setter
        public Builder minimum(Double minimum) {
            if (minimum == null) {
              throw new MissingRequiredPropertyException("GetDatabasePostgresqlConfigSharedBuffersPercentage", "minimum");
            }
            this.minimum = minimum;
            return this;
        }
        @CustomType.Setter
        public Builder requiresRestart(Boolean requiresRestart) {
            if (requiresRestart == null) {
              throw new MissingRequiredPropertyException("GetDatabasePostgresqlConfigSharedBuffersPercentage", "requiresRestart");
            }
            this.requiresRestart = requiresRestart;
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            if (type == null) {
              throw new MissingRequiredPropertyException("GetDatabasePostgresqlConfigSharedBuffersPercentage", "type");
            }
            this.type = type;
            return this;
        }
        public GetDatabasePostgresqlConfigSharedBuffersPercentage build() {
            final var _resultValue = new GetDatabasePostgresqlConfigSharedBuffersPercentage();
            _resultValue.description = description;
            _resultValue.example = example;
            _resultValue.maximum = maximum;
            _resultValue.minimum = minimum;
            _resultValue.requiresRestart = requiresRestart;
            _resultValue.type = type;
            return _resultValue;
        }
    }
}
